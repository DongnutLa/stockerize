import { useEffect, useState } from "react";
import { AuthUser, Product } from "../../../models";
import { getFromLocalStorage } from "../../../utils/functions";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { ProductsService } from "../../../services";
import ProductsTable from "../components/ProductsTable";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const productService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser?.token)

function ProductsAvailability() {
    const [products, setProducts] = useState<Product[]>([])
    const [isFetching, setIsFetching] = useState(false)

    useEffect(() => {
        setIsFetching(true);

        (async () => {
            const productsData = await productService.getProductsAvailability()
            setProducts(productsData.map(d => ({...d, key: d.id})))
        })()

        setIsFetching(false)
    }, [setProducts])
    
    return <ProductsTable
        products={products}
        isFetching={isFetching}
        handlePaginationChange={() => ({})}
        handleRowAction={() => ({ onClick: () => ({}) })}
        hiddenTabs={["prices"]}
    />;
}

export default ProductsAvailability;