import { Col, Row, Table, Tag } from "antd";
import { useCallback, useMemo } from "react";
import { Metadata, Order, PaymentMethod } from "../../../models";
import { getPageSizeOptions, numberToCurrency } from "../../../utils/functions";
import dayjs from "../../../utils/functions/dayjs";
import useWindowDimensions from "../../../utils/hooks/useWindowDimensions";

interface OrdersTableProps {
    orders: Order[]
    isFetching: boolean
    pagination?: Metadata
    handlePaginationChange: (page: number, pageSize: number) => void;
    handleRowAction: (record: any, index?: number) => { onClick: () => void };
}

function OrdersTable({orders, isFetching, pagination, handlePaginationChange, handleRowAction}: OrdersTableProps) {
  const { width } = useWindowDimensions()

  const renderExpandedRow = useCallback((order: Order) => {
    return (
      <Row gutter={[16, 16]}>
          <Col span={24}>
            Total: {numberToCurrency(order.totals.total)}
          </Col>
          <Col>
            Método: {<Tag color="red">{order.paymentMethod}</Tag>}
          </Col>
          <Col span={24}>
            Creación: {dayjs(order.createdAt).utc().format("D MMM, YYYY h:mm A")}
          </Col>
      </Row>
    )
  }, [])  
  
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
            responsive: ["md" as any],
            render: (total: number) => numberToCurrency(total)
          },
          {
            title: "Método",
            dataIndex: "paymentMethod",
            key: "paymentMethod",
            responsive: ["md" as any],
            render: (paymentMethod: PaymentMethod) => <Tag color="red">{paymentMethod}</Tag>
          },
          {
            title: "Creación",
            dataIndex: "createdAt",
            key: "createdAt",
            responsive: ["md" as any],
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
        expandable={{
          rowExpandable: () => width <= 576,
          defaultExpandAllRows: true,
          expandedRowRender: renderExpandedRow,
        }}
    /> );
}

export default OrdersTable;