// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

import { computed, reactive } from 'vue';
import { defineStore } from 'pinia';

import {
    AccountBalance,
    Coupon,
    CreditCard,
    DateRange,
    NativePaymentHistoryItem,
    PaymentsApi,
    PaymentsHistoryItem,
    PaymentsHistoryItemStatus,
    PaymentsHistoryItemType,
    ProjectCharges,
    ProjectUsagePriceModel,
    Wallet,
} from '@/types/payments';
import { PaymentsHttpApi } from '@/api/payments';

export class PaymentsState {
    public balance: AccountBalance = new AccountBalance();
    public creditCards: CreditCard[] = [];
    public paymentsHistory: PaymentsHistoryItem[] = [];
    public nativePaymentsHistory: NativePaymentHistoryItem[] = [];
    public projectCharges: ProjectCharges = new ProjectCharges();
    public usagePriceModel: ProjectUsagePriceModel = new ProjectUsagePriceModel();
    public startDate: Date = new Date();
    public endDate: Date = new Date();
    public coupon: Coupon | null = null;
    public wallet: Wallet = new Wallet();
}

export const useBillingStore = defineStore('billing', () => {
    const state = reactive<PaymentsState>(new PaymentsState());

    const api: PaymentsApi = new PaymentsHttpApi();

    async function getBalance(): Promise<AccountBalance> {
        const balance: AccountBalance = await api.getBalance();

        state.balance = balance;

        return balance;
    }

    async function getWallet(): Promise<void> {
        state.wallet = await api.getWallet();
    }

    async function claimWallet(): Promise<void> {
        state.wallet = await api.claimWallet();
    }

    async function setupAccount(): Promise<string> {
        return await api.setupAccount();
    }

    async function getCreditCards(): Promise<CreditCard[]> {
        const creditCards = await api.listCreditCards();

        state.creditCards = creditCards;

        return creditCards;
    }

    async function addCreditCard(token: string): Promise<void> {
        await api.addCreditCard(token);
    }

    function toggleCardSelection(id: string): void {
        state.creditCards = state.creditCards.map(card => {
            if (card.id === id) {
                card.isSelected = !card.isSelected;

                return card;
            }

            card.isSelected = false;

            return card;
        });
    }

    function clearCardsSelection(): void {
        state.creditCards = state.creditCards.map(card => {
            card.isSelected = false;

            return card;
        });
    }

    async function makeCardDefault(id: string): Promise<void> {
        await api.makeCreditCardDefault(id);

        state.creditCards = state.creditCards.map(card => {
            if (card.id === id) {
                card.isDefault = !card.isDefault;

                return card;
            }

            card.isDefault = false;

            return card;
        });
    }

    async function removeCreditCard(cardId: string): Promise<void> {
        await api.removeCreditCard(cardId);

        state.creditCards = state.creditCards.filter(card => card.id !== cardId);
    }

    async function getPaymentsHistory(): Promise<void> {
        state.paymentsHistory = await api.paymentsHistory();
    }

    async function getNativePaymentsHistory(): Promise<void> {
        state.nativePaymentsHistory = await api.nativePaymentsHistory();
    }

    async function getProjectUsageAndChargesCurrentRollup(): Promise<void> {
        const now = new Date();
        const endUTC = new Date(Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate(), now.getUTCHours(), now.getUTCMinutes()));
        const startUTC = new Date(Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), 1, 0, 0));

        state.projectCharges = await api.projectsUsageAndCharges(startUTC, endUTC);

        const dateRange = new DateRange(startUTC, endUTC);
        state.startDate = dateRange.startDate;
        state.endDate = dateRange.endDate;
    }

    async function getProjectUsageAndChargesPreviousRollup(): Promise<void> {
        const now = new Date();
        const startUTC = new Date(Date.UTC(now.getUTCFullYear(), now.getUTCMonth() - 1, 1, 0, 0));
        const endUTC = new Date(Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), 0, 23, 59, 59));

        state.projectCharges = await api.projectsUsageAndCharges(startUTC, endUTC);

        const dateRange = new DateRange(startUTC, endUTC);
        state.startDate = dateRange.startDate;
        state.endDate = dateRange.endDate;
    }

    async function getProjectUsagePriceModel(): Promise<void> {
        state.usagePriceModel = await api.projectUsagePriceModel();
    }

    async function applyCouponCode(code: string): Promise<void> {
        state.coupon = await api.applyCouponCode(code);
    }

    async function getCoupon(): Promise<void> {
        state.coupon = await api.getCoupon();
    }

    async function purchasePricingPackage(token: string): Promise<void> {
        await api.purchasePricingPackage(token);
    }

    function clear(): void {
        state.balance = new AccountBalance();
        state.creditCards = [];
        state.paymentsHistory = [];
        state.nativePaymentsHistory = [];
        state.projectCharges = new ProjectCharges();
        state.usagePriceModel = new ProjectUsagePriceModel();
        state.startDate = new Date();
        state.endDate = new Date();
        state.coupon = null;
        state.wallet = new Wallet();
    }

    const canUserCreateFirstProject = computed((): boolean => {
        return state.balance.sum > 0 || state.creditCards.length > 0;
    });

    const isTransactionProcessing = computed((): boolean => {
        return state.balance.sum === 0 &&
            state.paymentsHistory.some((item: PaymentsHistoryItem) => {
                return item.amount >= 50 && item.type === PaymentsHistoryItemType.Transaction &&
                    (
                        item.status === PaymentsHistoryItemStatus.Pending ||
                        item.status === PaymentsHistoryItemStatus.Paid ||
                        item.status === PaymentsHistoryItemStatus.Completed
                    );
            });
    });

    const isBalancePositive = computed((): boolean => {
        return state.balance.sum > 0;
    });

    return {
        state,
        canUserCreateFirstProject,
        isTransactionProcessing,
        isBalancePositive,
        getBalance,
        getWallet,
        claimWallet,
        setupAccount,
        getCreditCards,
        addCreditCard,
        toggleCardSelection,
        clearCardsSelection,
        makeCardDefault,
        removeCreditCard,
        getPaymentsHistory,
        getNativePaymentsHistory,
        getProjectUsageAndChargesCurrentRollup,
        getProjectUsageAndChargesPreviousRollup,
        getProjectUsagePriceModel,
        applyCouponCode,
        getCoupon,
        purchasePricingPackage,
        clear,
    };
});
