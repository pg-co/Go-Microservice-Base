const axios = require('axios')

async function main(){
    axios.get('http://localhost:8888/users').then((resp) => {
        console.log(resp)
    });
}

main()