import dayjs from "dayjs";
import es from "dayjs/locale/es"
import utc from "dayjs/plugin/utc";
dayjs.extend(utc).locale(es);

export default dayjs;