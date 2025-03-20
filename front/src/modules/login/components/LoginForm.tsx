import { useState } from 'react';
import { Form, Input, Button, Checkbox, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import 'antd/dist/reset.css'; // Asegúrate de importar los estilos de Ant Design
import Logo from '../../../Components/Logo';

const LoginForm = () => {
  const [loading, setLoading] = useState(false);

  const onFinish = (values: any) => {
    setLoading(true);
    // Simulando una llamada a una API
    setTimeout(() => {
      setLoading(false);
      message.success('Inicio de sesión exitoso');
      console.log('Valores recibidos:', values);
      // Aquí podrías redirigir al usuario a otra página
    }, 2000);
  };

  return (
    <div style={{ maxWidth: 300, margin: '0 auto', padding: '10%', border: '2px solid #FA8C16', borderRadius: '8px', backgroundColor: '#fff' }}>
        <Logo />
      <Form
        name="login_form"
        initialValues={{ remember: true }}
        onFinish={onFinish}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: 'Por favor ingresa tu usuario!' }]}
        >
          <Input prefix={<UserOutlined />} placeholder="Usuario" />
        </Form.Item>

        <Form.Item
          name="password"
          rules={[{ required: true, message: 'Por favor ingresa tu contraseña!' }]}
        >
          <Input.Password prefix={<LockOutlined />} placeholder="Contraseña" />
        </Form.Item>

        <Form.Item name="remember" valuePropName="checked">
          <Checkbox>Recordarme</Checkbox>
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading} block>
            Iniciar Sesión
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default LoginForm;