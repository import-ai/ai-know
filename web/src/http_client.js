export class HTTPError extends Error {
  constructor(status, data) {
    super('HTTP request failed with status ' + status)
    this.status = status
    this.data = data
  }
}

export class HTTPClient {
  constructor(baseURL) {
    this.baseURL = baseURL;
  }

  async request(endpoint, method, data, headers) {
    const config = {
      method,
      headers: {
        ...headers,
      },
    };

    if (data) {
      config.headers['Content-Type'] = 'application/json';
      config.body = JSON.stringify(data);
    }

    const resp = await fetch(`${this.baseURL}${endpoint}`, config);
    const resp_json = await resp.json()
    if (!resp.ok) {
      throw new HTTPError(resp.status, resp_json)
    }
    return resp_json
  }

  async get(endpoint, headers = {}) {
    return await this.request(endpoint, 'GET', null, headers);
  }

  async post(endpoint, data, headers = {}) {
    return await this.request(endpoint, 'POST', data, headers);
  }

  async put(endpoint, data, headers = {}) {
    return await this.request(endpoint, 'PUT', data, headers);
  }

  async delete(endpoint, headers = {}) {
    return await this.request(endpoint, 'DELETE', null, headers);
  }
}
