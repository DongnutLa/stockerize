import { useCallback, useEffect, useState } from "react";
import ProductsTable from "../components/ProductsTable";
import { AuthUser, Metadata, Product } from "../../../models";
import { ProductsService } from "../../../services";
import { getFromLocalStorage } from "../../../utils/functions";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { useNavigate } from "react-router";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const productService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser.token)

function ProductsList() {
    const [products, setProducts] = useState<Product[]>([])
    const [pagination, setPagination] = useState<Metadata>()
    const [params, setParams] = useState<{page: number, pageSize: number, search?: string}>({page: 1, pageSize: 20})
    const [isFetching, setIsFetching] = useState(false)

    const navigate = useNavigate()

    useEffect(() => {
        setIsFetching(true);

        (async () => {
            const productsData = await productService.getProductsList(params)
            setProducts(productsData.data.map(d => ({...d, key: d.id})))
            setPagination(productsData.metadata)
        })()

        setIsFetching(false)
    }, [params, setProducts])

    const handlePaginationChange = useCallback((page: number, pageSize: number) => {
        setParams({page, pageSize})
    }, [])

    const handleRowAction = (record: Product, _?: number) => ({
        onClick: () => navigate(`/products/${record.id}`),
      });

    return ( <>
        <ProductsTable products={products} pagination={pagination} isFetching={isFetching} handlePaginationChange={handlePaginationChange} handleRowAction={handleRowAction}/>
    </> );
}

export default ProductsList;