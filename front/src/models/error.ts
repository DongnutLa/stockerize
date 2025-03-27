export class ApiError {
  code: string;
  statusCode: number;
  message: string;
  detail?: any;

  constructor({
    code,
    statusCode,
    message,
    detail,
  }: {
    code: string;
    statusCode: number;
    message: string;
    detail?: any;
  }) {
    this.code = code;
    this.statusCode = statusCode;
    this.message = message;
    this.detail = detail;
  }
}
