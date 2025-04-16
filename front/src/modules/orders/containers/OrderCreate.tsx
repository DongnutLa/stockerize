import { useEffect } from "react";
import { AuthUser, emptyOrderDto } from "../../../models";
import { OrdersService, ProductsService } from "../../../services";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { getFromLocalStorage } from "../../../utils/functions";
import OrderForm from "../components/OrderForm";
import useOrderForm from "../hooks/useOrderForm";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const ordersService = new OrdersService(import.meta.env.VITE_API_BASE_URL, authUser?.token)
const productsService = new ProductsService(import.meta.env.VITE_API_BASE_URL, authUser.token)

function OrderCreate() {
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
    } = useOrderForm({type: "CREATE", ordersService, productsService})

    useEffect(() => {
        form.setFieldsValue(emptyOrderDto)
    }, [])

    return (
        <OrderForm
            type="CREATE"
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
        />
    );
}

export default OrderCreate;