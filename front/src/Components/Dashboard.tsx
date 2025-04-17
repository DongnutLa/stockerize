import './Dashboard.css';
import { AuthUser, OrderSummary, PAYMENT_METHOD_LABEL, PaymentMethod, SUMMARY_TYPE_TEXT, SummaryDash, SummaryType } from '../models';
import { getFromLocalStorage, numberToCurrency } from '../utils/functions';
import { AUTH_USER_KEY } from '../utils/constants';
import { OrdersService } from '../services';
import { useEffect, useState } from 'react';
import { Typography } from 'antd';
import dayjs from '../utils/functions/dayjs';

const { Title } = Typography;

const authUser = JSON.parse(getFromLocalStorage(AUTH_USER_KEY) as string) as AuthUser
const ordersService = new OrdersService(import.meta.env.VITE_API_BASE_URL, authUser?.token)

const SummaryDashboard = () => {
    const [isFetching, setIsFetching] = useState<boolean>(false)
    const [summaries, setSummaries] = useState<SummaryDash>()

    useEffect(() => {
        setIsFetching(true);

        (async () => {
            const summaryData = await ordersService.getSummary()
            setSummaries(summaryData)
        })()

        setIsFetching(false)
    }, [setIsFetching])

  const renderRange = (type: SummaryType, item: OrderSummary): string => {
    switch (type) {
        case SummaryType.DAILY:
            return dayjs(item.start).utc().format("D MMMM, YYYY")
        case SummaryType.WEEKLY:
            return `${dayjs(item.start).utc().format("D MMMM")} - ${dayjs(item.end).utc().format("D MMMM")}`
        case SummaryType.MONTHLY:
            return dayjs(item.start).utc().format("MMMM")
        default:
            return "";
    }
  }

  if (isFetching || !summaries) {
    return <p>Loading...</p>
  }

  return (
    <div className="dashboard">

      <div className='titles'>
        <Title level={4}>Compras</Title>
        <Title level={4}>Ventas</Title>
      </div>

      <div className='sections'>
      {Object.entries(summaries).map(([orderType, periodData]) => (
        <div key={orderType} className={`section ${orderType.toLowerCase()}`}>

          <div className="cards-container">
            {Object.entries(periodData).map(([summaryType, item]) => (
              <div key={`${summaryType}-${item.id}`} className="card">
              <div className="card-header">
                <div className="period">{SUMMARY_TYPE_TEXT[summaryType as SummaryType]}</div>
                <div className="date-range">
                    {renderRange(summaryType as SummaryType, item)}
                </div>
              </div>

              <div className="totals">
                <div className="total-item">
                  <div className="total-label">Transacciones</div>
                  <div className="total-value">{item.count}</div>
                </div>
                <div className="total-item">
                  <div className="total-label">Total</div>
                  <div className="total-value">{numberToCurrency(item.total)}</div>
                </div>
              </div>

              {item.paymentMethodDetails?.length && (
                <div className="payment-methods">
                  <h3>Métodos de Pago</h3>
                  {item.paymentMethodDetails.map((method) => {
                    const percentage = (method.total / item.total) * 100;
                    return (
                      <div key={method.paymentMethod} className="method-item">
                        <div className="method-name">{PAYMENT_METHOD_LABEL[method.paymentMethod as PaymentMethod]}</div>
                        <div className="method-stats">
                          <span>{method.count} ({numberToCurrency(method.total)})</span>
                          <span>{percentage.toFixed(1)}%</span>
                        </div>
                        <div className="progress-bar">
                          <div 
                            className="progress" 
                            style={{ width: `${percentage}%` }}
                          ></div>
                        </div>
                      </div>
                    );
                  })}
                </div>
              )}
            </div>
            ))}
          </div>
        </div>
      ))}
      </div>
    </div>
  );
};

export default SummaryDashboard;