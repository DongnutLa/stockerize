import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { getFromLocalStorage } from "../../../utils/functions";
import { AuthUser, Product } from "../../../models";
import { ProductsService } from "../../../services/products";
import ProductForm from "../components/ProductForm";
import useProductForm from "../hooks/useProductForm";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const productService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser?.token)

const type = "UPDATE"

function ProductUpdate() {
    const { productId } = useParams();

    const [product, setProduct] = useState<Product>()
    const [isFetching, setIsFetching] = useState(true)
    const {form, values, formLoading, onSubmitProduct, onStockUpdate, stockFormLoading} = useProductForm({product, productService, type })

    useEffect(() => {
        if (!productId) {
            setIsFetching(false)
            // navigate a 404
            return
        };

        setIsFetching(true)
        try {
            (async () => {
                const productData = await productService.getProductById(productId)
                form.setFieldsValue(productData);
                setProduct(productData);
            })()
        } catch (error) {
            // navigate a 404
        }

        setIsFetching(false)

        return () => {
            form.resetFields();
        };
    }, [productId, form]);

    if (isFetching) return <div>Loading...</div>

    return (
        <ProductForm
            type={type}
            product={product}
            form={form}
            values={values}
            formLoading={formLoading}
            onSubmitProduct={onSubmitProduct}
            onStockUpdate={onStockUpdate}
            stockFormLoading={stockFormLoading}
        />
    );
}

export default ProductUpdate;