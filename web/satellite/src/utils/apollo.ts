// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { ApolloClient, ApolloLink, HttpLink, InMemoryCache, ServerError } from '@apollo/client/core';
import { setContext } from '@apollo/client/link/context';
import { onError } from '@apollo/client/link/error';

import { AuthHttpApi } from '@/api/auth';
import { AnalyticsErrorEventSource } from '@/utils/constants/analyticsEventNames';
import { useNotificationsStore } from '@/store/modules/notificationsStore';

/**
 * Satellite url.
 */
const satelliteUrl = new HttpLink({
    uri: process.env.VUE_APP_ENDPOINT_URL,
});

/**
 * Adding additional headers.
 */
const authLink = setContext((_, { headers }) => {
    // return the headers to the context so httpLink can read them
    return {
        headers: {
            ...headers,
        },
    };
});

/**
 * Handling unauthorized error.
 */
const errorLink = onError(({ graphQLErrors, networkError }) => {
    const notificationsStore = useNotificationsStore();

    if (graphQLErrors?.length) {
        notificationsStore.notifyError({ message: graphQLErrors.join('\n'), source: AnalyticsErrorEventSource.OVERALL_GRAPHQL_ERROR });
    }

    if (networkError) {
        const nError = (networkError as ServerError);
        if (nError.statusCode === 401) {
            new AuthHttpApi().logout();
            notificationsStore.notifyError({ message: 'Session token expired', source: AnalyticsErrorEventSource.OVERALL_SESSION_EXPIRED_ERROR });
            setTimeout(() => {
                window.location.href = window.location.origin + '/login';
            }, 2000);
        } else {
            const error = typeof nError.result === 'string' ? nError.result : nError.result.error;
            nError.result && notificationsStore.notifyError({ message: error, source: AnalyticsErrorEventSource.OVERALL_GRAPHQL_ERROR });
        }
    }

    if (!(networkError && (networkError as ServerError).statusCode === 401)) {
        return;
    }
});

/**
 * Combining error and satellite urls.
 */
const link = ApolloLink.from([
    errorLink,
    authLink.concat(satelliteUrl),
]);

/**
 * Creating apollo client.
 */
export const apollo = new ApolloClient({
    link: link,
    cache: new InMemoryCache(),
    connectToDevTools: true,
});
