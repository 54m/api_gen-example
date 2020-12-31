import { APIClient, MockOption } from "./client_generated/api_client";

// the simplest
(async () => {
    const client = new APIClient();

    const resp = await client.getHealthCheck({});

    console.log(resp);
})();

// with options
(async () => {
    const client = new APIClient(
        "pass", // [optional] token for Authorization: Bearer
        {
            "X-Foobar": "foobar", // [optional] custom headers
        },
        "http://localhost:8888",  // [optional] custom endpoint
        {
            cache: "no-cache", // [optional] options for fetch API
        },
    );

    const resp = await client.api.putUser(
        {
            name: "54m",
            age: 100,
            gender: 1,
        },
        {
            "X-Foobar": "hoge", // [optional] custom headers
        },
        {
            mode: "cors" // [optional] options for fetch API
        },
    );

    console.log(resp);
})();

// mock mode
(async () => {
    const client = new APIClient();
    const mockOption: MockOption = {
        wait_ms: 10,
        target_file: 'default'
    }

    const resp = await client.api.user.getSearch({
        name: "54m",
        age: 100,
        gender: 1,
    }, undefined, {
        'mock_option': mockOption,
    });

    console.log(resp);
})();