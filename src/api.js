const baseURL = 'http://localhost:8080/api';



export default {
    getConfigs,
    updateConfigStatus,
    getChangeInfos,
    ping
};
async function getConfigs() {
    const url = `${baseURL}/configs`;
    return await fetch(url).then(res => res.json());
}


async function updateConfigStatus(id, status) {
    const url = `${baseURL}/config/status`;
    return await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ id, status })
    }).then(res => res.json());
}

async function getChangeInfos() {
    const url = `${baseURL}/changeInfos`;
    return await fetch(url).then(res => res.json());
}


async function ping() {
    const url = `${baseURL}/ping`;
    return await fetch(url).then(res => res.json());
}