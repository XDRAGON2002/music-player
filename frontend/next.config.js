/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
      domains: ["static.vecteezy.com", "images.unsplash.com", "cdn.vectorstock.com", "i.scdn.co"]
  }
}

module.exports = nextConfig