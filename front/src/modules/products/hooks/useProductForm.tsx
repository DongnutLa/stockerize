import { useForm, useWatch } from "antd/es/form/Form";
import { useNavigate } from "react-router";
import { ApiError, Product, ProductDTO, ProductStockDTO } from "../../../models";
import { useState } from "react";
import { ProductsService } from "../../../services/products";
import { ROUTES } from "../../../utils/constants";
import { message } from "antd";
import { sleep } from "../../../utils/functions";
import { ErrorProducts } from "../../../utils/errors";

interface UseProductFormProps {
    product?: Product
    productService: ProductsService
    type: "CREATE" | "UPDATE"
}

function useProductForm({product, productService, type}: UseProductFormProps) {
    const navigate = useNavigate()
    const [formLoading, setFormLoading] = useState(false)
    const [stockFormLoading, setStockFormLoading] = useState(false)

    const [form] = useForm<ProductDTO>();
    const values = useWatch<ProductDTO>([], form);

    const onSubmitProduct = async (values: ProductDTO) => {
        setFormLoading(true)
        try {
            if (type === "CREATE") {
                await productService.createProduct(values)
            } else {
                await productService.updateProduct(values)
            }
            navigate(ROUTES.products)
            message.success({content: `Producto |${values.name}| ${type === "CREATE" ? "creado" : "actualizado"} correctamente`, duration: 2})
        } catch (error) {
            const errMsg = ErrorProducts[(error as ApiError)?.code]
            let content = `Hubo un error al ${type === "CREATE" ? "crear" : "actualizar"} el producto |${values.name}|`
            if (errMsg) {
                content = errMsg
            }

            message.error({content, duration: 2})
        }

        setFormLoading(false)
    }

    const onStockUpdate = async (values: {stock: number, cost: number}) => {
        if (!product) return;

        setStockFormLoading(true)

        let updateType: ProductStockDTO["updateType"] = "INFO"
        let quantity = product.stockSummary?.available ?? 0
    
        if (quantity !== values.stock) {
            updateType = values.stock < quantity ? "DECREASE" : "INCREASE"
            quantity = Math.abs(quantity - values.stock)
        }

        const stockDTO: ProductStockDTO = {
            id: product.id,
            updateType: updateType,
            cost: values.cost,
            quantity: updateType === "INFO" ? 0 : quantity,
            unitPrice: {
                subUnit: 1,
            },
        }

        try {
            await productService.updateStock(stockDTO)
            await sleep()
            navigate(ROUTES.products)
            message.success({content: `Inventario del producto |${product.name}| actualizado correctamente`, duration: 2})
        } catch (error) {
            const errMsg = ErrorProducts[(error as ApiError)?.code]
            let content = `Hubo un error al actualizar el inventario del producto |${product.name}|`
            if (errMsg) {
                content = errMsg
            }

            message.error({content, duration: 2})
        }

        setStockFormLoading(false)
    }

    return {
        form,
        values,
        formLoading,
        onSubmitProduct,
        onStockUpdate,
        stockFormLoading,
    };
}

export default useProductForm;