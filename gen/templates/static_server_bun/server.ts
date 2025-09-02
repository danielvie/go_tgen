// This script sets up a simple Bun server to serve all static files
// from the './static' directory. It uses Bun's built-in APIs,
// so no 'npm install' is needed.

import { existsSync } from 'fs';

// The Bun.serve function creates a new HTTP server.
const server = Bun.serve({
  // The port the server will listen on.
  port: 3000,

  // The fetch handler processes incoming requests.
  async fetch(request) {
    const url = new URL(request.url);

    // If the path is the root '/', serve the index.html file.
    let filePath = `${import.meta.dir}/static${url.pathname}`;
    if (url.pathname === '/') {
      filePath = `${import.meta.dir}/static/index.html`;
    }

    // Check if the requested file exists in the static directory.
    // We use `existsSync` for a quick check.
    if (existsSync(filePath)) {
      // If the file exists, return it using `Bun.file()`.
      // Bun will automatically handle the correct Content-Type.
      const file = Bun.file(filePath);
      return new Response(file);
    } else {
      // If the file does not exist, return a 404 Not Found response.
      return new Response("Not Found", { status: 404 });
    }
  },
});

console.log(`Bun server is running on http://localhost:${server.port}`);
console.log("Serving all files from the `./static` directory.");