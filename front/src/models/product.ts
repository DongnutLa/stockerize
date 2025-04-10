export type Product = {
    id: string
    name: string
    sku: string
    prices: Price[]
    unit: ProductUnit
    createdAt: string
    updatedAt: string
    stockSummary: StockSummary
  }
  
export type Price = {
    subUnit: number
    price: number
}

export type StockSummary = {
id: string
cost: number
quantity: number
available: number
sold: number
}

export type ProductDTO = {
    id?: string
    name: string
    sku: string
    prices: Price[]
    unit: ProductUnit
    stock?: {
        cost?: number
        quantity?: number
    }
}
export type ProductStockDTO = {
    id: string
    updateType: "INCREASE" | "DECREASE" | "INFO"
    cost: number
    quantity: number
    unitPrice: {
        subUnit: number
    }
}

export enum ProductUnit {
    PC = "PC",
    LT = "LT",
    KG = "KG",
    LB = "LB",
}

export const PRODUCT_UNIT_NAME: Record<ProductUnit, string> = {
    PC: "Unidades",
    LT: "Litros",
    KG: "Kilos",
    LB: "Libras",
}

export const productUnitOptions = Object.entries(PRODUCT_UNIT_NAME).map(([value, label]) => ({
    value,
    label
}))

export function subUnitText(subunit:number) {
    switch (subunit) {
        case 0.25:
            return "¼"
        case 0.5:
            return "½"
        case 0.75:
            return "¾"
        case 1:
            return "1"
        default:
            return ""
    }
}

export function subUnitOptions(unit: ProductUnit) {
    switch (unit) {
        case ProductUnit.LB:
            return [{label: "¼", value: 0.25},{label: "½", value: 0.5},{label: "¾", value: 0.75},{label: "1", value: 1},{label: "2 (1kg)", value: 2}]
        case ProductUnit.KG:
            return [{label: "¼", value: 0.25},{label: "½", value: 0.5},{label: "¾", value: 0.75},{label: "1", value: 1}]
        case ProductUnit.LT:
            return [{label: "¼", value: 0.25},{label: "½", value: 0.5},{label: "¾", value: 0.75},{label: "1", value: 1}]
        case ProductUnit.PC:
        default:
            return [{label: "1", value: 1}]
    }
}