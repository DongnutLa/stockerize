import { OrderProducts, OrderType } from "../../models"

export const calculateProductSubtotal = (product: OrderProducts, orderType: OrderType): number => {
    const price = orderType === "SALE" ? product.price : product.cost
    
    const subtotal = price * product.quantity
    return subtotal
}

export const getMaxToSell = (product: OrderProducts, products: OrderProducts[], ordType: OrderType): number | undefined => {
  if (ordType === "PURCHASE") return undefined;

  const variants = products.filter(p => p.id === product.id && product.unitPrice.subUnit !== p.unitPrice.subUnit)
  const selectedQty = variants.reduce((acc, curr) => acc += (curr.unitPrice.subUnit * curr.quantity), 0)
  
  const max = Math.floor(((product.available ?? 0) - selectedQty) / product.unitPrice.subUnit)
  return max
}