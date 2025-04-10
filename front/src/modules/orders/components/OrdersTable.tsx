import { Table, Tag } from "antd";
import { useMemo } from "react";
import { Metadata, Order, PaymentMethod } from "../../../models";
import { getPageSizeOptions, numberToCurrency } from "../../../utils/functions";
import dayjs from "../../../utils/functions/dayjs";

interface OrdersTableProps {
    orders: Order[]
    isFetching: boolean
    pagination?: Metadata
    handlePaginationChange: (page: number, pageSize: number) => void;
    handleRowAction: (record: any, index?: number) => { onClick: () => void };
}

function OrdersTable({orders, isFetching, pagination, handlePaginationChange, handleRowAction}: OrdersTableProps) {
    const columns = useMemo(() => {
        return [
          {
            title: "Número",
            dataIndex: "consecutive",
            key: "consecutive",
          },
          {
            title: "Total",
            dataIndex: ["totals", "total"],
            key: "total",
            render: (total: number) => numberToCurrency(total)
          },
          {
            title: "Método",
            dataIndex: "paymentMethod",
            key: "paymentMethod",
            render: (paymentMethod: PaymentMethod) => <Tag color="red">{paymentMethod}</Tag>
          },
          {
            title: "Creación",
            dataIndex: "createdAt",
            key: "createdAt",
            render: (createdAt: string) => dayjs(createdAt).utc().format("D MMM, YYYY h:mm A"),
          },
        ];
      }, []);

    return ( <Table
        dataSource={orders}
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

export default OrdersTable;