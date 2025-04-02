export function numberToCurrency(
    amount: number,
    currency = 'USD',
    locale = 'en-US',
    options = {}
  ) {
    const defaultOptions = {
      style: 'currency',
      currency,
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
      ...options // Allow overriding defaults
    };
  
    return new Intl.NumberFormat(locale, defaultOptions as any).format(amount);
  }