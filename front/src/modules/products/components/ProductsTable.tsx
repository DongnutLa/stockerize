import { Space, Table, Tag } from "antd";
import { useMemo } from "react";
import { Metadata, Price, Product, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import { getPageSizeOptions, numberToCurrency } from "../../../utils/functions";

interface ProductsTableProps {
    products: Product[]
    isFetching: boolean
    pagination?: Metadata
    handlePaginationChange: (page: number, pageSize: number) => void;
    handleRowAction: (record: any, index?: number) => { onClick: () => void };
}

function ProductsTable({products, isFetching, pagination, handlePaginationChange, handleRowAction}: ProductsTableProps) {
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
            title: "Precios",
            dataIndex: "prices",
            key: "prices",
            render: (prices: Price[]) => <Space direction="vertical">
                {prices.map(price => (
                    <div><Tag color="magenta">{subUnitText(price.subUnit)}</Tag> <span>{numberToCurrency(price.price)}</span></div>
                ))}
            </Space>
          },
          {
            title: "Disponible",
            dataIndex: ["stockSummary", "available"],
            key: "stock",
            render: (stock: number, reg: Product) => `${stock} ${PRODUCT_UNIT_NAME[reg.unit]}`
          }/* ,
          {
            title: "Creación",
            dataIndex: "createdAt",
            key: "createdAt",
            render: (createdAt: string) => dayjs(createdAt).utc().format("D MMM, YYYY h:mm A"),
          }, */
        ];
      }, []);

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
    /> );
}

export default ProductsTable;