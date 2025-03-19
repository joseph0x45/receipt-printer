export type BillItemT = {
  product_id: string
  product_name: string
  quantity: number
  total_price: number
}

export type Bill = {
  client_name: string
  issued_at: string
  items: Array<BillItemT>
}
