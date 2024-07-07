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

  async logout() {
    await this.http_client.post('/logout', {})
  }

  async listKbs() {
    return await this.http_client.get('/kbs')
  }

  async createKb(kb) {
    return await this.http_client.post('/kbs', kb)
  }

  async listNotes(kb_id) {
    return await this.http_client.get(`/kbs/${kb_id}/notes`)
  }

  async createNote(kb_id, note) {
    return await this.http_client.post(`/kbs/${kb_id}/notes`, note)
  }

  async updateNote(kb_id, note) {
    return await this.http_client.put(`/kbs/${kb_id}/notes/${note.id}`, note)
  }
}

export const api_client = new APIClient('/api')
