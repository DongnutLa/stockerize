import { useParams } from "react-router";
import useOrderForm from "../hooks/useOrderForm";
import { calculateProductSubtotal, getFromLocalStorage } from "../../../utils/functions";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { AuthUser, Order } from "../../../models";
import { OrdersService, ProductsService } from "../../../services";
import OrderForm from "../components/OrderForm";
import { useEffect, useState } from "react";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const ordersService = new OrdersService(import.meta.env.VITE_API_BASE_URL, authUser?.token)
const productsService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser?.token)

function OrderUpdate() {
    const { orderId } = useParams();

    const [order, setOrder] = useState<Order>()
    const [isFetching, setIsFetching] = useState(true)
    const {
        form,
        formLoading,
        onSubmitOrder,
        values,
        searchProducts,
        productOptions,
        setProductsOptions,
        handleSelectOption,
        onChangeQuantity,
        onChangePrice,
    } = useOrderForm({type: "UPDATE", ordersService, productsService, order})

    console.log(values)

    useEffect(() => {
            if (!orderId) {
                setIsFetching(false)
                // navigate a 404
                return
            };
    
            setIsFetching(true)
            try {
                (async () => {
                    const productData = await ordersService.getOrderById(orderId)

                    productData.products = productData.products.map(prd => {
                        prd.subtotal = calculateProductSubtotal(prd, productData.type)
                        
                        return prd
                    })

                    form.setFieldsValue(productData);
                    setOrder(productData);
                })()
            } catch (error) {
                // navigate a 404
            }
    
            setIsFetching(false)
    
            return () => {
                form.resetFields();
            };
        }, [orderId, form]);
    
        if (isFetching) return <div>Loading...</div>

    return (
        <OrderForm
            type="UPDATE"
            form={form}
            formLoading={formLoading}
            onSubmitOrder={onSubmitOrder}
            values={values}
            searchProducts={searchProducts}
            productOptions={productOptions}
            setProductsOptions={setProductsOptions}
            handleSelectOption={handleSelectOption}
            onChangeQuantity={onChangeQuantity}
            onChangePrice={onChangePrice}
            disabled
        />
    );
}

export default OrderUpdate;