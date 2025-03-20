import { useParams } from "react-router";

function ProductUpdate() {
    let params = useParams();
    console.log(params.productId)

    return ( <div>Hola Product update</div> );
}

export default ProductUpdate;