import { Col, Row, Space, Table, Tag } from "antd";
import { useCallback, useMemo } from "react";
import { Metadata, Price, Product, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import { getPageSizeOptions, numberToCurrency } from "../../../utils/functions";
import useWindowDimensions from "../../../utils/hooks/useWindowDimensions";

interface ProductsTableProps {
    products: Product[]
    isFetching: boolean
    pagination?: Metadata
    handlePaginationChange: (page: number, pageSize: number) => void;
    handleRowAction: (record: any, index?: number) => { onClick: () => void };
    hiddenTabs?: string[]
}

function ProductsTable({
  products,
  isFetching,
  pagination,
  hiddenTabs,
  handlePaginationChange,
  handleRowAction
}: ProductsTableProps) {
  const { width } = useWindowDimensions()

  const renderExpandedRow = useCallback((product: Product) => {
    return (
      <Row gutter={[16, 16]}>
          <Col span={24}>
            Código: {product.sku}
          </Col>
          <Col>
            Unidad: {PRODUCT_UNIT_NAME[product.unit]}
          </Col>
          <Col span={24}>
            {(product.prices ?? []).map(price => (
              <div><Tag color="magenta">{subUnitText(price.subUnit)}</Tag> <span>{numberToCurrency(price.price)}</span></div>
            ))}
          </Col>
          <Col span={24}>
            Disponible: {`${product.stockSummary.available} ${PRODUCT_UNIT_NAME[product.unit]}`}
          </Col>
      </Row>
    )
  }, [])

    const columns = useMemo(() => {
        let cols = [
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
            title: "Precios",
            dataIndex: "prices",
            key: "prices",
            responsive: ["md" as any],
            render: (prices: Price[]) => <Space direction="vertical">
                {(prices ?? []).map(price => (
                    <div><Tag color="magenta">{subUnitText(price.subUnit)}</Tag> <span>{numberToCurrency(price.price)}</span></div>
                ))}
            </Space>
          },
          {
            title: "Disponible",
            dataIndex: ["stockSummary", "available"],
            key: "stock",
            responsive: ["md" as any],
            render: (stock: number, reg: Product) => `${stock} ${PRODUCT_UNIT_NAME[reg.unit]}`
          }
        ];

        if (hiddenTabs?.length) {
          cols = cols.filter(tab => !hiddenTabs.includes(tab.key))
        }

        return cols
      }, [hiddenTabs]);

    return ( <Table
        dataSource={products}
        columns={columns}
        loading={isFetching}
        tableLayout="fixed"
        scroll={{ y: 450 }}
        pagination={{
            current: pagination?.page ?? 1,
            pageSize: pagination?.pageSize ?? 20,
            total: pagination?.count ?? 0,
            pageSizeOptions: getPageSizeOptions(pagination?.count ?? 0),
            showSizeChanger: true,
            onChange: handlePaginationChange,
        }}
        onRow={handleRowAction}
        expandable={{
          rowExpandable: () => width <= 576,
          defaultExpandAllRows: true,
          expandedRowRender: renderExpandedRow,
        }}
    /> );
}

export default ProductsTable;