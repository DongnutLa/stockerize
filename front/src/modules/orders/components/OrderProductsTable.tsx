import { InputNumber, Table, Tag } from "antd";
import { useCallback } from "react";
import { OrderProducts, OrderType, Price, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import { getMaxToSell, numberToCurrency } from "../../../utils/functions";

interface OrderProductsTableProps {
    products: OrderProducts[]
    orderType: OrderType
    onChangeQuantity: (productId: string, subUnit: number, quantity: number) => void
    onChangePrice: (productId: string, subUnit: number, quantity: number) => void
}

function OrderProductsTable({products, orderType, onChangeQuantity, onChangePrice}: OrderProductsTableProps) {
    const columns = useCallback((products: OrderProducts[]) => {
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
                <InputNumber defaultValue={qty} min={1} max={getMaxToSell(product, products, orderType)} value={product.quantity} onChange={(value) => onChangeQuantity(product.id, product.unitPrice.subUnit, value ?? 1)} />
            )
          },
          {
            title: orderType === "SALE" ? "Precio unitario" : "Costo unitario",
            dataIndex: orderType === "SALE" ? "price" : "cost",
            key: orderType === "SALE" ? "price" : "cost",
            render: (qty: number, product: OrderProducts) => (
                <InputNumber defaultValue={qty} value={orderType === "SALE" ? product.price : product.cost} onChange={(value) => onChangePrice(product.id, product.unitPrice.subUnit, value ?? 1)} />
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
        columns={columns(products)}
        tableLayout="fixed"
        scroll={{ y: 450 }}
        pagination={false}
    />
    );
}

export default OrderProductsTable;