import { Button, Col, Divider, Form, FormInstance, Input, InputNumber, Row, Select } from "antd";
import { Product, PRODUCT_UNIT_NAME, ProductDTO, productUnitOptions, subUnitOptions } from "../../../models";

import { DeleteOutlined } from "@ant-design/icons";
import Title from "antd/es/typography/Title";

interface ProductFormProps {
    type: "CREATE" | "UPDATE"
    product?: Product
    form: FormInstance<ProductDTO>;
    values: ProductDTO
    formLoading: boolean
    onStockUpdate: (values: {stock: number, cost: number}) => void
    onSubmitProduct: (values: ProductDTO) => void
    stockFormLoading: boolean
}

function ProductForm({ type, product, form, values, formLoading, onStockUpdate, stockFormLoading, onSubmitProduct }: ProductFormProps) {
    return ( 
        <>
        <Title>{type === "CREATE" ? "Crear" : "Actualizar"} producto</Title>
        <Divider>Catálogo</Divider>
        <Form
            form={form}
            style={{ maxWidth: 1000, paddingTop: 24, paddingRight: 24, overflow: 'auto', height: 'calc(100% - 155px)' }}
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            onFinish={onSubmitProduct}
        >
        <Row>
            <Col xs={24} sm={24} span={12}>
                <Form.Item<string>
                    label="Nombre"
                    name="name"
                    rules={[{ required: true, message: 'Agrega un nombre!' }]}
                >
                    <Input />
                </Form.Item>
            </Col>
            <Col xs={24} sm={24} span={12}>
                <Form.Item<string>
                    label="Código"
                    name="sku"
                    rules={[{ required: true, message: 'Agrega un código!' }]}
                >
                    <Input />
                </Form.Item>
            </Col>

            <Col xs={24} sm={24} span={12}>
                <Form.Item<string>
                    label="Unidad"
                    name="unit"
                    rules={[{ required: true, message: 'Selecciona una unidad!' }]}
                >
                    <Select
                        placeholder="Selecciona una unidad"
                        optionFilterProp="unit"
                        options={productUnitOptions}
                    />
                </Form.Item>
            </Col>

            <Col span={24}>
                <Form.Item required label="Precios" labelCol={{ span: 4}} wrapperCol={{ span: 20 }}>
                <Form.List name="prices">
                    {(fields, {add, remove}) => (
                        <>
                            {fields.map((field) => (
                                <Row key={field.name} gutter={[16, 16]}>
                                    <Col xs={6} sm={6} span={10}>
                                        <Form.Item<string>
                                            name={[field.name, "subUnit"]}
                                            rules={[{ required: true, message: "Agrega una unidad!" }]}
                                        >
                                            <Select options={subUnitOptions(values?.unit)} />
                                        </Form.Item>
                                    </Col>

                                    <Col xs={16} sm={16} span={12}>
                                        <Form.Item<string>
                                            key={field.name}
                                            name={[field.name, "price"]}
                                            rules={[{ required: true, message: "Agrega un precio!" }]}
                                        >
                                            <InputNumber<number>
                                                prefix="$"
                                                formatter={(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                                                parser={(value) => value?.replace(/\$\s?|(,*)/g, '') as unknown as number}
                                                width="100%"
                                                style={{width: "100%"}}
                                                min={1}
                                            />
                                        </Form.Item>
                                    </Col>

                                    <Col span={2}>
                                        <Button
                                            type="dashed"
                                            danger
                                            icon={<DeleteOutlined />}
                                            onClick={() => remove(field.name)}
                                        />
                                    </Col>
                                </Row>
                            ))}
                            <Button type="dashed" onClick={() => add()} disabled={fields.length >= subUnitOptions(values?.unit).length}>
                                Agregar precio
                            </Button>
                        </>
                    )}
                </Form.List>
                </Form.Item>

                {type === "CREATE" && (
                    <>
                        <Col xs={24} sm={24} span={18}>
                        <Form.Item<number>
                            name={["stock", "quantity"]}
                            rules={[{ required: true, message: "Agrega una cantidad!" }]}
                            label="Inventario inicial"
                        >
                            <InputNumber<number>
                                width="100%"
                                style={{width: "100%"}}
                                min={0}
                                suffix={PRODUCT_UNIT_NAME[values?.unit]}
                            />
                        </Form.Item>
                        </Col>

                        <Col xs={24} sm={24} span={18}>
                        <Form.Item<number>
                            name={["stock", "cost"]}
                            rules={[{ required: true, message: "Agrega un costo!" }]}
                            label="Costo unitario"
                        >
                            <InputNumber<number>
                                prefix="$"
                                formatter={(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                                parser={(value) => value?.replace(/\$\s?|(,*)/g, '') as unknown as number}
                                width="100%"
                                style={{width: "100%"}}
                                min={1}
                            />
                        </Form.Item>
                        </Col>
                    </>
                )}

                <Form.Item hidden name="id" />

                <Col span={24}>
                    <Form.Item style={{ float: "right" }}>
                        <Button
                            type="primary"
                            htmlType="submit"
                            loading={formLoading}
                            disabled={formLoading}
                        >
                            {type === "CREATE" ? "Crear" : "Actualizar"} producto
                        </Button>
                    </Form.Item>
                </Col>
            </Col>
        </Row>
        </Form>

        {!!product && type === "UPDATE" && (
            <>
                <Divider>Inventario</Divider>
                <Form
                    initialValues={{
                        stock: product.stockSummary?.available ?? 0, 
                        cost: product.stockSummary?.cost ?? 0,
                    }}
                    style={{ maxWidth: 1000, paddingTop: 24, paddingRight: 24 }}
                    labelCol={{ span: 8 }}
                    wrapperCol={{ span: 16 }}
                    onFinish={onStockUpdate}
                >
                    <Row>
                        <Col span={12}>
                            <Form.Item required label="Inventario" name="stock">
                                <InputNumber<number>
                                    width="100%"
                                    style={{width: "100%"}}
                                    min={0}
                                    suffix={PRODUCT_UNIT_NAME[values?.unit ?? product.unit]}
                                />
                            </Form.Item>
                        </Col>

                        <Col span={12}>
                            <Form.Item required label="Costo" name="cost">
                                <InputNumber<number>
                                    prefix="$"
                                    formatter={(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                                    parser={(value) => value?.replace(/\$\s?|(,*)/g, '') as unknown as number}
                                    width="100%"
                                    style={{width: "100%"}}
                                    min={1}
                                />
                            </Form.Item>
                        </Col>
                    </Row>

                    <Col span={24}>
                    <Form.Item style={{ float: "right" }}>
                        <Button
                            type="primary"
                            htmlType="submit"
                            loading={stockFormLoading}
                            disabled={stockFormLoading}
                        >
                            Actualizar inventario
                        </Button>
                    </Form.Item>
                </Col>
                </Form>
            </>
        )}
        </>
    );
}

export default ProductForm;