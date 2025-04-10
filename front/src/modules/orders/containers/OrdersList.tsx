import { useCallback, useEffect, useState } from "react";
import { AuthUser, Metadata, Order } from "../../../models";
import { OrderParams, OrdersService } from "../../../services";
import { AUTH_USER_KEY } from "../../../utils/constants";
import { getFromLocalStorage } from "../../../utils/functions";
import { useNavigate } from "react-router";
import { Tabs } from "antd";
import {CaretUpOutlined, CaretDownOutlined} from '@ant-design/icons';
import OrdersTable from "../components/OrdersTable";

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const orderService = new OrdersService(import.meta.env.VITE_API_BASE_URL, authUser.token)

function OrdersList() {
    const [sales, setSales] = useState<Order[]>([])
    const [salesPagination, setSalesPagination] = useState<Metadata>()
    const [salesParams, setSalesParams] = useState<OrderParams>({page: 1, pageSize: 20, orderType: "SALE"})
    const [isFetchingSales, setIsFetchingSales] = useState(false)
    
    const [purchases, setPurchases] = useState<Order[]>([])
    const [purchasesPagination, setPurchasesPagination] = useState<Metadata>()
    const [purchasesParams, setPurchasesParams] = useState<OrderParams>({page: 1, pageSize: 20, orderType: "PURCHASE"})
    const [isFetchingPurchases, setIsFetchingPurchases] = useState(false)

    const navigate = useNavigate()

    // Sales
    useEffect(() => {
        setIsFetchingSales(true);

        (async () => {
            const productsData = await orderService.getOrdersList(salesParams)
            setSales(productsData.data.map(d => ({...d, key: d.id})))
            setSalesPagination(productsData.metadata)
        })()

        setIsFetchingSales(false)
    }, [salesParams, setSales, setSalesPagination, setIsFetchingPurchases])

    const handleSalesPaginationChange = useCallback((page: number, pageSize: number) => {
        setSalesParams({...salesParams, page, pageSize})
    }, [])

    // Purchases
    useEffect(() => {
        setIsFetchingPurchases(true);

        (async () => {
            const productsData = await orderService.getOrdersList(purchasesParams)
            setPurchases(productsData.data.map(d => ({...d, key: d.id})))
            setPurchasesPagination(productsData.metadata)
        })()

        setIsFetchingPurchases(false)
    }, [purchasesParams, setPurchases, setPurchasesPagination, setIsFetchingPurchases])

    const handlePurchasesPaginationChange = useCallback((page: number, pageSize: number) => {
        setPurchasesParams({...purchasesParams, page, pageSize})
    }, [])

    // both
    const handleRowAction = (record: Order, _?: number) => ({
        onClick: () => navigate(`/orders/${record.id}`),
      });

    return (
        <Tabs
            defaultActiveKey="sales"
            centered
            items={[
                {
                    key: "sales",
                    label: "Ventas",
                    children: <OrdersTable
                                orders={sales}
                                pagination={salesPagination}
                                isFetching={isFetchingSales}
                                handlePaginationChange={handleSalesPaginationChange}
                                handleRowAction={handleRowAction}
                            />,
                    icon: <CaretUpOutlined />,
                },
                {
                    key: "purchases",
                    label: "Compras",
                    children: <OrdersTable
                                orders={purchases}
                                pagination={purchasesPagination}
                                isFetching={isFetchingPurchases}
                                handlePaginationChange={handlePurchasesPaginationChange}
                                handleRowAction={handleRowAction}
                            />,
                    icon: <CaretDownOutlined />,
                },
            ]}
        />
    );
}

export default OrdersList;