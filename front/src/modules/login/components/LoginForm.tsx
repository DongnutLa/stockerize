import { useState } from 'react';
import { Form, Input, Button, Checkbox, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import 'antd/dist/reset.css'; // Asegúrate de importar los estilos de Ant Design
import Logo from '../../../Components/Logo';
import { AuthService } from '../../../services';
import { saveToLocalStorage } from '../../../utils/functions';
import { AUTH_USER_KEY, ROUTES } from '../../../utils/constants';
import { useNavigate } from 'react-router';

const authService = new AuthService(import.meta.env.VITE_API_BASE_URL)

const LoginForm = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const onFinish = async (values: any) => {
    setLoading(true);

    try {
      const authUser = await authService.login(values)
      saveToLocalStorage(AUTH_USER_KEY, JSON.stringify(authUser))
      navigate(ROUTES.root)
      message.success({content: `Bienvenid@ ${authUser.name}!`, duration: 2})
    } catch (error) {
      message.error({content: "Credenciales inválidas", duration: 2})
    }

    setLoading(false);
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
          name="email"
          rules={[{ required: true, message: 'Por favor ingresa tu correo!' }]}
        >
          <Input type='email' prefix={<UserOutlined />} placeholder="Correo" />
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