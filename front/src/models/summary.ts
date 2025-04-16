import { OrderType } from "./order";

export interface OrderSummary {
    id: string
    orderType: string
    summaryType: string
    paymentMethod: string
    count: number
    total: number
    start: string
    end: string
    updatedAt: string
    paymentMethodDetails: PaymentMethodDetail[]
}
export interface PaymentMethodDetail {
    paymentMethod: string
    count: number
    total: number
}


export enum SummaryType {
    DAILY = "DAILY",
    MONTHLY = "MONTHLY",
    WEEKLY = "WEEKLY"
}

export type SummaryDash = Record<OrderType, Record<SummaryType, OrderSummary>>

export const SUMMARY_TYPE_TEXT = {
    [SummaryType.DAILY]: "Diario",
    [SummaryType.WEEKLY]: "Semanal",
    [SummaryType.MONTHLY]: "Mensual",
}