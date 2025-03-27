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

export function subUnitText(subunit:number) {
    switch (subunit) {
        case 0.25:
            return "1/4"
        case 0.5:
            return "1/2"
        case 0.75:
            return "3/4"
        case 1:
            return "1"
        default:
            return ""
    }
}