import { Col, InputNumber, Row, Table, Tag } from "antd";
import { useCallback } from "react";
import { OrderProducts, OrderType, Price, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import { getMaxToSell, numberToCurrency } from "../../../utils/functions";
import useWindowDimensions from "../../../utils/hooks/useWindowDimensions";

interface OrderProductsTableProps {
    products: OrderProducts[]
    orderType: OrderType
    onChangeQuantity: (productId: string, subUnit: number, quantity: number) => void
    onChangePrice: (productId: string, subUnit: number, quantity: number) => void
}

function OrderProductsTable({products, orderType, onChangeQuantity, onChangePrice}: OrderProductsTableProps) {
  const { width } = useWindowDimensions()
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
            responsive: ["md" as any],
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
            responsive: ["md" as any],
            render: (qty: number, product: OrderProducts) => (
                <InputNumber defaultValue={qty} min={1} max={getMaxToSell(product, products, orderType)} value={product.quantity} onChange={(value) => onChangeQuantity(product.id, product.unitPrice.subUnit, value ?? 1)} />
            )
          },
          {
            title: orderType === "SALE" ? "Precio unitario" : "Costo unitario",
            dataIndex: orderType === "SALE" ? "price" : "cost",
            key: orderType === "SALE" ? "price" : "cost",
            responsive: ["md" as any],
            render: (qty: number, product: OrderProducts) => (
                <InputNumber defaultValue={qty} value={orderType === "SALE" ? product.price : product.cost} onChange={(value) => onChangePrice(product.id, product.unitPrice.subUnit, value ?? 1)} />
            )
          },
          {
            title: "Subtotal",
            dataIndex: "subtotal",
            key: "subtotal",
            responsive: ["md" as any],
            render: (subtotal: number) => numberToCurrency(subtotal)
          }
        ];
      }, [products, orderType, onChangeQuantity, onChangePrice]);

  const renderExpandedRow = useCallback((record: OrderProducts) => {
    
    return (
      <Row gutter={[16, 16]}>
        <Col span={24}>
          Código: {record.sku}
        </Col>
        <Col span={24}>
          Cantidad: <InputNumber defaultValue={record.quantity} min={1} max={getMaxToSell(record, products, orderType)} value={record.quantity} onChange={(value) => onChangeQuantity(record.id, record.unitPrice.subUnit, value ?? 1)} />
        </Col>
        <Col span={24}>
          {orderType === "SALE" ? "Precio unitario: " : "Costo unitario: "} <InputNumber defaultValue={orderType === "SALE" ? record.price : record.cost} value={orderType === "SALE" ? record.price : record.cost} onChange={(value) => onChangePrice(record.id, record.unitPrice.subUnit, value ?? 1)} />
        </Col>
        <Col span={24}>
          Subtotal: {numberToCurrency(record.subtotal ?? 0)}
        </Col>
      </Row>
    )
  }, [numberToCurrency, getMaxToSell, onChangeQuantity, onChangePrice])

  return (
    <Table
      dataSource={products.map(p => ({...p, key: `${p.id}|${p.unitPrice.subUnit}`}))}
      columns={columns(products)}
      tableLayout="fixed"
      scroll={{ y: 450 }}
      pagination={false}
      expandable={{
        rowExpandable: () => width <= 576,
        defaultExpandAllRows: true,
        expandedRowRender: renderExpandedRow,
      }}
    />
  );
}

export default OrderProductsTable;