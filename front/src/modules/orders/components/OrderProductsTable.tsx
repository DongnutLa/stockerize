import { InputNumber, Table, Tag } from "antd";
import { useMemo } from "react";
import { OrderProducts, OrderType, Price, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import { numberToCurrency } from "../../../utils/functions";

interface OrderProductsTableProps {
    products: OrderProducts[]
    orderType: OrderType
    onChangeQuantity: (productId: string, quantity: number) => void
    onChangePrice: (productId: string, quantity: number) => void
}

function OrderProductsTable({products, orderType, onChangeQuantity, onChangePrice}: OrderProductsTableProps) {
    const columns = useMemo(() => {
        return [
          {
            title: "Nombre",
            dataIndex: "name",
            key: "name",
          },
          {
            title: "Código",
            dataIndex: "sku",
            key: "sku",
          },
          {
            title: "Unidad",
            dataIndex: "unitPrices",
            key: "unitPrices",
            render: (_: Price[], product: OrderProducts) => <Tag color="magenta">{subUnitText(product.unitPrice.subUnit)} {PRODUCT_UNIT_NAME[product.unit]}</Tag>
          },
          {
            title: "Cantidad",
            dataIndex: "quantity",
            key: "quantity",
            render: (qty: number, product: OrderProducts) => (
                <InputNumber defaultValue={qty} value={product.quantity} onChange={(value) => onChangeQuantity(product.id, value ?? 1)} />
            )
          },
          {
            title: orderType === "SALE" ? "Precio unitario" : "Costo unitario",
            dataIndex: orderType === "SALE" ? "price" : "cost",
            key: orderType === "SALE" ? "price" : "cost",
            render: (qty: number, product: OrderProducts) => (
                <InputNumber defaultValue={qty} value={orderType === "SALE" ? product.price : product.cost} onChange={(value) => onChangePrice(product.id, value ?? 1)} />
            )
          },
          {
            title: "Subtotal",
            dataIndex: "subtotal",
            key: "subtotal",
            render: (subtotal: number) => numberToCurrency(subtotal)
          }
        ];
      }, [products, orderType, onChangeQuantity, onChangePrice]);
    return (
        <Table
        dataSource={products.map(p => ({...p, key: `${p.id}|${p.unitPrice.subUnit}`}))}
        columns={columns}
        tableLayout="fixed"
        scroll={{ y: 450 }}
        pagination={false}
    />
    );
}

export default OrderProductsTable;