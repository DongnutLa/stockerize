import { Button, Col, Form, FormInstance, InputNumber, Radio, Row, Tag, Typography } from "antd";
import { emptyOrderDto, Order, OrderDTO, PAYMENT_METHOD_LABEL, PaymentMethod, Product, PRODUCT_UNIT_NAME, subUnitText } from "../../../models";
import Title from "antd/es/typography/Title";
import DebouncedSearchSelect from "./SearchSelect";
import OrderProductsTable from "./OrderProductsTable";
import { DataNode } from "antd/es/tree";
import { numberToCurrency } from "../../../utils/functions";
import { useMemo } from "react";
import Select, { BaseOptionType } from "antd/es/select";

interface OrderFormProps {
    type: "CREATE" | "UPDATE"
    order?: Order
    disabled?: boolean
    form: FormInstance<OrderDTO>
    values: OrderDTO
    formLoading: boolean
    onSubmitOrder: (values: OrderDTO) => void
    searchProducts: (value: string) => Promise<Product[]>
    productOptions: Product[]
    setProductsOptions: (products: Product[]) => void
    handleSelectOption: (value: string) => void
    onChangeQuantity: (productId: string, quantity: number) => void
    onChangePrice: (productId: string, quantity: number) => void
}

function OrderForm({
    type,
    form,
    values,
    disabled,
    formLoading,
    onSubmitOrder,
    searchProducts,
    productOptions,
    setProductsOptions,
    handleSelectOption,
    onChangeQuantity,
    onChangePrice,
}: OrderFormProps) {

    const selectOptions = useMemo(() => {
        if (values?.type === "SALE") {
            return productOptions.map(p => ({
                key: `${p.id}|1|parent`,
                value: `${p.id}|1|parent`,
                title: (
                    <Row>
                        <Col span={12}>
                            {p.name}
                        </Col>
                        
                        <Col span={12}>
                            Disponible: <Tag color="orange">{p.stockSummary?.available ?? 0}</Tag> {PRODUCT_UNIT_NAME[p.unit]}
                        </Col>
                    </Row>
                ),
                disabled: (values?.products ?? []).some(e => `${e.id}|1` === `${p.id}|1`),
                children: p.prices.map(pr => ({
                    key: `${p.id}|${pr.subUnit}|child`,
                    value: `${p.id}|${pr.subUnit}|child`,
                    title: (
                        <Row>
                            <Col>
                                <Tag color="orange">{subUnitText(pr.subUnit)} {PRODUCT_UNIT_NAME[p.unit]}</Tag>
                            </Col>
                            <Col>
                                <span>{numberToCurrency(pr.price)}</span>
                            </Col>
                        </Row>
                    ),
                    disabled: (values?.products ?? []).some(e => `${e.id}|${e.unitPrice.subUnit}` === `${p.id}|${pr.subUnit}`),
                })),
            })) as DataNode[]
        }

        return productOptions.map(p => ({
            key: `${p.id}|1`,
            value: `${p.id}|1`,
            disabled: (values?.products ?? []).some(e => e.id === p.id),
            label: (
                <Row>
                    <Col span={13}>
                        {p.name}
                    </Col>
                    
                    <Col span={3}>
                        <Tag color="orange">{p.stockSummary?.available ?? 0}</Tag>
                    </Col>

                    <Col span={4}>
                        {PRODUCT_UNIT_NAME[p.unit]}
                    </Col>

                    <Col span={4}>
                        {numberToCurrency(p.stockSummary.cost)}
                    </Col>
                </Row>
            ),
        })) as BaseOptionType[]
    }, [values?.type, productOptions, values?.products])

    return (
        <>
            <Title>{type === "CREATE" ? "Crear" : "Actualizar"} orden</Title>
            <Form
                form={form}
                layout="vertical"
                style={{
                    maxWidth: 1000,
                    maxHeight: "calc(100% - 137px)",
                    padding: 24,
                    overflowY: "auto",
                    backgroundColor: "white",
                    borderRadius: 12,
                    margin: 12,
                }}
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 24 }}
                // initialValues={emptyOrderDto}
                onFinish={onSubmitOrder}
                disabled={disabled}
            >
                <Row gutter={24}>
                    {type === "UPDATE" && values?.consecutive ? (
                        <Col span={24}>
                            <Form.Item labelAlign="left" layout="horizontal" name="consecutive" label="Número de orden" wrapperCol={{style: {textAlign: "left"}}}>
                                <Tag color="gold">{values.consecutive}</Tag>
                            </Form.Item>
                        </Col>
                    ) : <></>}

                    <Col span={24}>
                        <Form.Item<string>
                            label="Tipo de orden"
                            name="type"
                            rules={[{ required: true, message: 'Selecciona un tipo de orden!' }]}
                        >
                            <Radio.Group
                                block
                                options={[
                                    {label: "Orden de venta", value: "SALE"},
                                    {label: "Orden de compra", value: "PURCHASE"}
                                ]}
                                defaultValue="SALE"
                                optionType="button"
                                buttonStyle="solid"
                                />
                        </Form.Item>

                        <Form.Item label="Productos" name="products" rules={[{ required: true, message: "Selecciona al menos un producto"}]}>
                            <DebouncedSearchSelect<Product>
                                placeholder="Buscar productos"
                                valueProp="id"
                                fetchFn={searchProducts}
                                options={selectOptions}
                                setOptions={setProductsOptions}
                                handleSelectOption={handleSelectOption}
                                excludedOptions={[]}
                                isTreeSelect={values?.type === "SALE"}
                            />

                            <OrderProductsTable
                                onChangeQuantity={onChangeQuantity}
                                onChangePrice={onChangePrice}
                                products={values?.products ?? []}
                                orderType={values?.type}
                            />
                        </Form.Item>
                    </Col>

                    <Col span={12}>
                        <Form.Item label="Descuentos" name="discount">
                            <InputNumber<number>
                                prefix="$"
                                formatter={(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                                parser={(value) => value?.replace(/\$\s?|(,*)/g, '') as unknown as number}
                                width="100%"
                                style={{width: "100%"}}
                                min={0}
                                max={values?.totals?.subtotal}
                            />
                        </Form.Item>
                    </Col>

                    <Col span={12}>
                        <Form.Item labelCol={{ style: { width: "100%" }}} label="Método de pago" name="paymentMethod" rules={[{ required: true, message: "Selecciona un método de pago"}]}>
                            <Select
                                placeholder="Seleccione método de pago"
                                filterOption={false}
                                style={{ width: "100%" }}
                                options={Object.values(PaymentMethod).map(p => ({
                                    label: <Tag color="blue">{PAYMENT_METHOD_LABEL[p]}</Tag>,
                                    value: p,
                                }))}
                            />
                        </Form.Item>
                    </Col>

                    <div style={{padding: 20, backgroundColor: "#fff5ea", borderRadius: 12, width: "50%"}}>
                        <Typography.Title style={{fontSize: "2rem"}}>Totales</Typography.Title>
                        <Col span={24}>
                            <Form.Item layout="horizontal" labelAlign="left" labelCol={{ style: {width: "50%"}}} wrapperCol={{ style: {textAlign: "left"}}} label="Subtotal" name={["totals", "subtotal"]}>
                                {numberToCurrency(values?.totals?.subtotal ?? 0)}
                            </Form.Item>
                        </Col>

                        <Col span={24}>
                            <Form.Item layout="horizontal" labelAlign="left" labelCol={{ style: {width: "50%"}}} wrapperCol={{ style: {textAlign: "left"}}} label="Descuento" name={["totals", "discount"]}>
                                {numberToCurrency(values?.totals?.discount ?? 0)}
                            </Form.Item>
                        </Col>

                        <Col span={24}>
                            <Form.Item layout="horizontal" labelAlign="left" labelCol={{ style: {width: "50%"}}} wrapperCol={{ style: {textAlign: "left"}}} label="Total" name={["totals", "total"]}>
                                {numberToCurrency(values?.totals?.total ?? 0)}
                            </Form.Item>
                        </Col>
                    </div>

                    <Form.Item hidden name="id" />
                    <Form.Item hidden name="consecutive" />

                    <Col span={24}>
                        <Form.Item style={{ float: "right" }}>
                            <Button
                                type="primary"
                                htmlType="submit"
                                loading={formLoading}
                                disabled={formLoading || disabled}
                            >
                                {type === "CREATE" ? "Crear" : "Actualizar"} orden
                            </Button>
                        </Form.Item>
                    </Col>
                </Row>
            </Form>
        </>
    );
}

export default OrderForm;