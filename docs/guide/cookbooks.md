# Cookbooks

Multi-step recipes and real-world patterns for the most-used gox packages. Each cookbook combines several APIs the way you would in a production service.

::: tip How to use these
Start with the package **Overview**, then read the **Cookbook** for end-to-end patterns, then dive into individual API pages for Node.js vs Go comparisons.
:::

## Available cookbooks

| Package | Node.js analog | Cookbook |
|---------|----------------|----------|
| [`cond`](/packages/cond/) | `? :`, `??` | [cond cookbook](/packages/cond/cookbook) |
| [`slice`](/packages/slice/) | lodash / Array.* | [slice cookbook](/packages/slice/cookbook) |
| [`maputil`](/packages/maputil/) | lodash objects | [maputil cookbook](/packages/maputil/cookbook) |
| [`http`](/packages/http/) | express | [http cookbook](/packages/http/cookbook) |
| [`client`](/packages/client/) | axios | [client cookbook](/packages/client/cookbook) |
| [`env`](/packages/env/) | dotenv | [env cookbook](/packages/env/cookbook) |
| [`async`](/packages/async/) | Promise.all | [async cookbook](/packages/async/cookbook) |
| [`validate`](/packages/validate/) | zod / joi | [validate cookbook](/packages/validate/cookbook) |
| [`db`](/packages/db/) | knex | [db cookbook](/packages/db/cookbook) |
| [`redis`](/packages/redis/) | ioredis | [redis cookbook](/packages/redis/cookbook) |
| [`mongo`](/packages/mongo/) | mongoose | [mongo cookbook](/packages/mongo/cookbook) |
| [`jwt`](/packages/jwt/) | jsonwebtoken | [jwt cookbook](/packages/jwt/cookbook) |

## Example: Express → gox/http

The [http cookbook](/packages/http/cookbook) walks through a full server with middleware, JSON binding, mounted routers, CORS, and typed errors — the same patterns as a typical Express app.

## Examples from tests

Individual API pages now include **Example from tests** sections where the gox test suite demonstrates real usage. These snippets are extracted automatically from `*_test.go` files in the library.
