export const dynamic = 'force-dynamic' // defaults to auto
export const dynamicParams = true

export async function POST(request: Request, { params }: { params: { name: string } }) {
  if (!params || !('name' in params)) {
    return Response.json({ 'status': '404' })
  }
  const j = await request.json()
  // console.log(j)
  const resp = await fetch(process.env.SERVER_ADDR + "/api/" + params.name, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(j),
  })
  return new Response(resp.body)
  // return await resp.text();
}