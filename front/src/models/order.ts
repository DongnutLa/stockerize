import { Price, ProductUnit } from "./product"
import { User } from "./user"

export type OrderType = "SALE" | "PURCHASE"
export enum PaymentMethod {
    CASH = "CASH",
    NEQUI = "NEQUI",
    DAVIPLATA = "DAVIPLATA"
}

export const PAYMENT_METHOD_LABEL: Record<PaymentMethod, string> = {
    [PaymentMethod.CASH]: "Efectivo",
    [PaymentMethod.NEQUI]: "Nequi",
    [PaymentMethod.DAVIPLATA]: "Daviplata",
}

export type OrderTotals = {
    subtotal: number
    discount: number
    total: number
}
export const ORDER_TOTALS_DEFAULT: OrderTotals = {
    subtotal: 0,
    discount: 0,
    total: 0,
}

export type OrderProducts = {
    id: string
    name: string
    sku: string
    quantity: number
    price: number
    cost: number
    unit: ProductUnit
    unitPrice: Omit<Price, "price">

    unitPrices?: Price[]
    subtotal?: number
}

export type Order = {
    id: string
    type: OrderType
    consecutive: string
    products: OrderProducts[]
    user: User
    totals: OrderTotals
    paymentMethod: PaymentMethod
    createdAt: string
    updatedAt: string
}

export type OrderDTO = {
    id?: string
    consecutive?: string
    type: OrderType
    products: OrderProducts[]
    totals: OrderTotals
    discount: number
    paymentMethod: PaymentMethod
}
export const emptyOrderDto: OrderDTO = {
    type: "SALE",
    products: [],
    totals: ORDER_TOTALS_DEFAULT,
    discount: 0,
    paymentMethod: PaymentMethod.CASH,
}