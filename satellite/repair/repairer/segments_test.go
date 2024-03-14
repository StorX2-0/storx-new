// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package repairer_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"storj.io/common/memory"
	"storj.io/common/pb"
	"storj.io/common/storj"
	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/storj/private/testplanet"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/buckets"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/overlay"
	"storj.io/storj/satellite/repair/queue"
)

func TestSegmentRepairPlacement(t *testing.T) {
	piecesCount := 4
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 8, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.ReconfigureRS(1, 2, piecesCount, piecesCount),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		require.NoError(t, planet.Uplinks[0].CreateBucket(ctx, planet.Satellites[0], "testbucket"))

		_, err := planet.Satellites[0].API.Buckets.Service.UpdateBucket(ctx, buckets.Bucket{
			ProjectID: planet.Uplinks[0].Projects[0].ID,
			Name:      "testbucket",
			Placement: storj.EU,
		})
		require.NoError(t, err)

		for _, node := range planet.StorageNodes {
			require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, node.ID(), "PL"))
		}

		err = planet.Uplinks[0].Upload(ctx, planet.Satellites[0], "testbucket", "object", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		type testCase struct {
			piecesOutOfPlacement int
			piecesAfterRepair    int
		}

		for _, tc := range []testCase{
			// all pieces/nodes are out of placement, repair download/upload should be triggered
			{piecesOutOfPlacement: piecesCount, piecesAfterRepair: piecesCount},
			// few pieces/nodes are out of placement, repair download/upload should be triggered
			{piecesOutOfPlacement: piecesCount - 2, piecesAfterRepair: piecesCount},
			// single piece/node is out of placement, NO download/upload repair, we are only removing piece from segment
			// as segment is still above repair threshold
			{piecesOutOfPlacement: 1, piecesAfterRepair: piecesCount - 1},
		} {
			for _, node := range planet.StorageNodes {
				require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, node.ID(), "PL"))
			}

			require.NoError(t, planet.Satellites[0].Repairer.Overlay.DownloadSelectionCache.Refresh(ctx))

			segments, err := planet.Satellites[0].Metabase.DB.TestingAllSegments(ctx)
			require.NoError(t, err)
			require.Len(t, segments, 1)
			require.Len(t, segments[0].Pieces, piecesCount)

			for _, piece := range segments[0].Pieces[:tc.piecesOutOfPlacement] {
				require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, piece.StorageNode, "US"))
			}

			// confirm that some pieces are out of placement
			ok, err := allPiecesInPlacement(ctx, planet.Satellites[0].Overlay.Service, segments[0].Pieces, segments[0].Placement)
			require.NoError(t, err)
			require.False(t, ok)

			require.NoError(t, planet.Satellites[0].Repairer.Overlay.DownloadSelectionCache.Refresh(ctx))

			_, err = planet.Satellites[0].Repairer.SegmentRepairer.Repair(ctx, &queue.InjuredSegment{
				StreamID: segments[0].StreamID,
				Position: segments[0].Position,
			})
			require.NoError(t, err)

			// confirm that all pieces have correct placement
			segments, err = planet.Satellites[0].Metabase.DB.TestingAllSegments(ctx)
			require.NoError(t, err)
			require.Len(t, segments, 1)
			require.NotNil(t, segments[0].RepairedAt)
			require.Len(t, segments[0].Pieces, tc.piecesAfterRepair)

			ok, err = allPiecesInPlacement(ctx, planet.Satellites[0].Overlay.Service, segments[0].Pieces, segments[0].Placement)
			require.NoError(t, err)
			require.True(t, ok)
		}
	})
}

func TestSegmentRepairPlacementAndClumped(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 8, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: testplanet.Combine(
				testplanet.ReconfigureRS(1, 2, 4, 4),
				func(log *zap.Logger, index int, config *satellite.Config) {
					config.Checker.DoDeclumping = true
					config.Repairer.DoDeclumping = true
				},
			),
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		require.NoError(t, planet.Uplinks[0].CreateBucket(ctx, planet.Satellites[0], "testbucket"))

		_, err := planet.Satellites[0].API.Buckets.Service.UpdateBucket(ctx, buckets.Bucket{
			ProjectID: planet.Uplinks[0].Projects[0].ID,
			Name:      "testbucket",
			Placement: storj.EU,
		})
		require.NoError(t, err)

		for _, node := range planet.StorageNodes {
			require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, node.ID(), "PL"))
		}

		err = planet.Uplinks[0].Upload(ctx, planet.Satellites[0], "testbucket", "object", testrand.Bytes(5*memory.KiB))
		require.NoError(t, err)

		for _, node := range planet.StorageNodes {
			require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, node.ID(), "PL"))
		}

		require.NoError(t, planet.Satellites[0].Repairer.Overlay.DownloadSelectionCache.Refresh(ctx))

		segments, err := planet.Satellites[0].Metabase.DB.TestingAllSegments(ctx)
		require.NoError(t, err)
		require.Len(t, segments, 1)
		require.Len(t, segments[0].Pieces, 4)

		// set nodes to the same placement/country and put all nodes into the same net to mark them as clumped
		node0 := planet.FindNode(segments[0].Pieces[0].StorageNode)
		for _, piece := range segments[0].Pieces {
			require.NoError(t, planet.Satellites[0].Overlay.Service.TestNodeCountryCode(ctx, piece.StorageNode, "US"))

			local := node0.Contact.Service.Local()
			checkInInfo := overlay.NodeCheckInInfo{
				NodeID:     piece.StorageNode,
				Address:    &pb.NodeAddress{Address: local.Address},
				LastIPPort: local.Address,
				LastNet:    node0.Contact.Service.Local().Address,
				IsUp:       true,
				Operator:   &local.Operator,
				Capacity:   &local.Capacity,
				Version:    &local.Version,
			}
			err = planet.Satellites[0].DB.OverlayCache().UpdateCheckIn(ctx, checkInInfo, time.Now().UTC(), overlay.NodeSelectionConfig{})
			require.NoError(t, err)
		}

		// confirm that some pieces are out of placement
		ok, err := allPiecesInPlacement(ctx, planet.Satellites[0].Overlay.Service, segments[0].Pieces, segments[0].Placement)
		require.NoError(t, err)
		require.False(t, ok)

		require.NoError(t, planet.Satellites[0].Repairer.Overlay.DownloadSelectionCache.Refresh(ctx))

		_, err = planet.Satellites[0].Repairer.SegmentRepairer.Repair(ctx, &queue.InjuredSegment{
			StreamID: segments[0].StreamID,
			Position: segments[0].Position,
		})
		require.NoError(t, err)

		// confirm that all pieces have correct placement
		segments, err = planet.Satellites[0].Metabase.DB.TestingAllSegments(ctx)
		require.NoError(t, err)
		require.Len(t, segments, 1)
		require.NotNil(t, segments[0].RepairedAt)
		require.Len(t, segments[0].Pieces, 4)

		ok, err = allPiecesInPlacement(ctx, planet.Satellites[0].Overlay.Service, segments[0].Pieces, segments[0].Placement)
		require.NoError(t, err)
		require.True(t, ok)
	})
}

func allPiecesInPlacement(ctx context.Context, overaly *overlay.Service, pieces metabase.Pieces, placement storj.PlacementConstraint) (bool, error) {
	for _, piece := range pieces {
		nodeDossier, err := overaly.Get(ctx, piece.StorageNode)
		if err != nil {
			return false, err
		}
		if !placement.AllowedCountry(nodeDossier.CountryCode) {
			return false, nil
		}
	}
	return true, nil
}
