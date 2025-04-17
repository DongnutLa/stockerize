import { useForm, useWatch } from "antd/es/form/Form";
import { useNavigate } from "react-router";
import { ApiError, Order, ORDER_TOTALS_DEFAULT, OrderDTO, OrderProducts, Product } from "../../../models";
import { useCallback, useEffect, useState } from "react";
import { ROUTES } from "../../../utils/constants";
import { message } from "antd";
import { ErrorOrders } from "../../../utils/errors";
import { OrdersService, ProductsService } from "../../../services";
import { cloneDeep } from "lodash";
import { calculateProductSubtotal } from "../../../utils/functions";

interface UseOrderFormProps {
    ordersService: OrdersService
    productsService: ProductsService
    type: "CREATE" | "UPDATE"
    order?: Order
}

const useOrderForm = ({ordersService, productsService, type}: UseOrderFormProps) => {
    const navigate = useNavigate()
    const [formLoading, setFormLoading] = useState(false)

    const [form] = useForm<OrderDTO>();
    const values = useWatch<OrderDTO>([], form);

    const validateStock = (values: OrderDTO): boolean => {
        const {type, products} = values;

        if (type === "PURCHASE") return true;

        const data: Record<string, {
            available: number // total
            quantity: number
            subUnit: number
        }[]> = {}
        products.forEach((variant) => {
            const available = variant.available ?? 0
            const quantity = variant.quantity
            const subUnit = variant.unitPrice.subUnit

            const subb = {
                available,
                quantity,
                subUnit,
            }

            if (variant.id in data) {
                data[variant.id] = data[variant.id].concat(subb)
            } else {
                data[variant.id] = [subb]
            }
        })

        const invalid = Object.values(data).some((units) => {
            const maxAvailable = units[0].available
            const total = units.reduce((acc, curr) => acc += (curr.subUnit * curr.quantity), 0)

            return maxAvailable < total
        })

        return !invalid
    }

    const onSubmitOrder = async (values: OrderDTO) => {
        try {
            await form.validateFields({recursive: true})
            const validStock = validateStock(values)
            if (!validStock) throw new Error("Invalid stock")
        } catch (error) {
            message.error({content: `Hubo un error al ${type === "CREATE" ? "crear" : "actualizar"} la orden. Verifica los datos de la orden!`, duration: 2})
            return
        }

        const dto = cloneDeep(values)
        dto.products = dto.products.map(p => {
            delete p.unitPrices
            delete p.subtotal
            delete p.available
            return p
        })

        setFormLoading(true)
        var number = ""
        try {
            if (type === "CREATE") {
                const res = await ordersService.createOrder(dto)
                number = res.consecutive
            } else {
                const res = await ordersService.updateOrder(dto)
                number = res.consecutive
            }
            navigate(ROUTES.orders)
            message.success({content: `Order |${number}| ${type === "CREATE" ? "creada" : "actualizada"} correctamente`, duration: 2})
        } catch (error) {
            const errMsg = ErrorOrders[(error as ApiError)?.code]
            let content = `Hubo un error al ${type === "CREATE" ? "crear" : "actualizar"} la orden |${number}|`
            if (errMsg) {
                content = errMsg
            }

            message.error({content, duration: 2})
        }

        setFormLoading(false)
    }

    const [productOptions, setOptions] = useState<Product[]>([])

    const searchProducts = async (value: string): Promise<Product[]> => {
        const res = await productsService.getProductsList({page: 1, pageSize: 50, search: value})
        return res.data
    }
    const setProductsOptions = (products: Product[]) => {
        setOptions(products)
    }

    const handleSelectOption = (value: string) => {
        const [productId, subUnit] = value.split("|")
        const selectedProduct = productOptions.find(p => p.id === productId)

        if (selectedProduct) {
            const orderProduct: OrderProducts = {
                id: selectedProduct.id,
                name: selectedProduct.name,
                sku: selectedProduct.sku,
                quantity: 1,
                price: selectedProduct.prices.find(p => p.subUnit === +subUnit)?.price ?? 0,
                cost: selectedProduct.stockSummary?.cost,
                unit: selectedProduct.unit,
                unitPrice: { subUnit: +subUnit },
                unitPrices: selectedProduct.prices,
                available: selectedProduct.stockSummary?.available ?? 0,
            }

            orderProduct.subtotal = calculateProductSubtotal(orderProduct, values?.type)

            const newProducts = [...(values?.products ?? []), orderProduct]
            form.setFieldValue("products", newProducts)
        }
    }

    useEffect(() => {
        const totals = (values?.products ?? []).reduce(
            (acc, curr) => {
                const subtotal = calculateProductSubtotal(curr, values?.type)

                return {...acc, subtotal: acc.subtotal + subtotal}
            },
            ORDER_TOTALS_DEFAULT,
          );

          totals.discount = values?.discount ?? 0
          totals.total = totals.subtotal - totals.discount

          form.setFieldValue("totals", totals)
    }, [values?.products, values?.discount])

    const onChangeQuantity = useCallback((productId: string, subUnit: number, quantity: number) => {
        const newProducts = cloneDeep(values?.products ?? [])

        const productIdx = newProducts.findIndex(p => p.id === productId && p.unitPrice.subUnit === subUnit)

        if (productIdx >= 0) {
            newProducts[productIdx].quantity = quantity
            
            const subtotal = calculateProductSubtotal(newProducts[productIdx], values?.type)
            newProducts[productIdx].subtotal = subtotal
        }


        form.setFieldValue("products", newProducts)
    }, [values, form])

    const onChangePrice = useCallback((productId: string, subUnit: number, price: number) => {
        const newProducts = cloneDeep(values?.products ?? [])

        const productIdx = newProducts.findIndex(p => p.id === productId && p.unitPrice.subUnit === subUnit)

        if (productIdx >= 0) {
            if (values?.type === "SALE") {
                newProducts[productIdx].price = price
            }
            if (values?.type === "PURCHASE") {
                newProducts[productIdx].cost = price
            }

            const subtotal = calculateProductSubtotal(newProducts[productIdx], values?.type)
            newProducts[productIdx].subtotal = subtotal
        }

        form.setFieldValue("products", newProducts)
    }, [values, form])

    return {
        form,
        values,
        formLoading,
        onSubmitOrder,
        searchProducts,
        productOptions,
        setProductsOptions,
        handleSelectOption,
        onChangeQuantity,
        onChangePrice,
    };
}

export default useOrderForm;