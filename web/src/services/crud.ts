import { Vue } from 'vue-property-decorator';

export class CrudService<T extends { id: string }> {
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  protected readonly baseUrl: string;

  async list(): Promise<T[]> {
    const res = await Vue.$api.get<T[]>(this.baseUrl);
    if (!res.success || !res.data) throw new Error(res.message);
    return res.data;
  }

  async get(id: string): Promise<T> {
    const res = await Vue.$api.get<T>(`${this.baseUrl}/${id}`);
    if (!res.success || !res.data) throw new Error(res.message);
    return res.data as T;
  }

  async save(item: T): Promise<T> {
    let res;
    if (!item.id) {
      res = await Vue.$api.post(this.baseUrl, { ...item, id: undefined });
    } else {
      res = await Vue.$api.put(`${this.baseUrl}/${item.id}`, { ...item });
    }
    if (res.status === 422) {
      throw new Error('Invalidated data');
    }
    if (!res.success || !res.data) {
      throw new Error(res.message);
    }
    return res.data as T;
  }

  async remove(id: string) {
    const res = await Vue.$api.delete(`${this.baseUrl}/${id}`);
    if (!res.success) {
      throw new Error(res.message);
    }
    return true;
  }
}
