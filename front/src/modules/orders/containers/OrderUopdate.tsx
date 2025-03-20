import { useParams } from "react-router";

function OrderUpdate() {
    let params = useParams();
    console.log(params.orderId)

    return ( <div>Hola order update</div> );
}

export default OrderUpdate;