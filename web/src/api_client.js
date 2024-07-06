import { HTTPClient } from '@/http_client'

export class APIClient {
  constructor(baseURL) {
    this.http_client = new HTTPClient(baseURL)
  }

  async login(name, password) {
    await this.http_client.post('/login', {
      name: name,
      password: password,
    })
  }

  async register(name, password) {
    await this.http_client.post('/register', {
      name: name,
      password: password,
    })
  }

  async getAuthorizedUser() {
    const resp = await this.http_client.get('/user')
    return resp.name
  }
}

export const api_client = new APIClient('http://localhost/api')
