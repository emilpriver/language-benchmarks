import { serve } from '@hono/node-server'
import { Hono } from 'hono'

const app = new Hono()
app.get('/json', (c) => c.json({ message: "Hello from node" }))

serve(app)
