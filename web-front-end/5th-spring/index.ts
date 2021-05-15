import { Application, send } from "https://deno.land/x/oak/mod.ts";

const app = new Application();

// handle login api
app.use(async (ctx, next) => {
    if (ctx.request.url.pathname === '/api/login') {
        const result = ctx.request.body({ type: "json"});
        const { username, password } = await result.value;
        if (username === 'admin' && password === 'password')
            ctx.response.body = { status: true, message: "Login Success", token: "fake_token" };
        else
            ctx.response.body = { status: false, message: "Login Failed" };
    } else {
        await next()
    }
});

// handle the index html of child pages 
app.use(async (ctx, next) => {
    if (["/app"].includes(ctx.request.url.pathname)) {
        await send(ctx, `${ctx.request.url.pathname}/index.html`, {
            root: `${Deno.cwd()}/static`,
        })
    } else {
        await next()
    }
});

// handle static contents
app.use(async ctx => {
    await send(ctx, ctx.request.url.pathname, {
        root: `${Deno.cwd()}/static`,
        index: "index/index.html"
    })
});

await app.listen({ port: 8000 });