# React App


## Steps

- create a new project
```powershell
npm create vite@laetst my-app
```
or

```powershell
bun create vite my-app
```

- add tailwind

install tailwind:

```powershell
npm install tailwindcss @tailwindcss/vite
```
or

```powershell
bun install tailwindcss @tailwindcss/vite
```

edit `vite.config.ts` with:
```ts
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
})
```

import the css:
```css
@import "tailwindcss";
```


## troubleshoot
if error: `Cannot find module '@babel/traverse/`
then
```powershell
npm install --save-dev @babel/traverse @babel/core
```