import { AuthUser } from "../../../models";
import { ProductsService } from "../../../services/products";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { getFromLocalStorage } from "../../../utils/functions";
import ProductForm from "../components/ProductForm";
import useProductForm from "../hooks/useProductForm";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const productService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser?.token)

const type = "CREATE"   

function ProductCreate() {

    const {
        form,
        values,
        formLoading,
        onSubmitProduct,
        onStockUpdate,
        stockFormLoading
    } = useProductForm({productService, type })
    
    return (
        <ProductForm
            type={type}
            form={form}
            values={values}
            formLoading={formLoading}
            onSubmitProduct={onSubmitProduct}
            onStockUpdate={onStockUpdate}
            stockFormLoading={stockFormLoading}
        />
    );
}

export default ProductCreate;