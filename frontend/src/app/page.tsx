'use client'

import { useState, useEffect } from 'react'
import { motion } from 'framer-motion'
import { MapPin, Camera, BookOpen, Globe } from 'lucide-react'
import Link from 'next/link'

export default function Home() {
  const [currentTime, setCurrentTime] = useState(new Date())

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date())
    }, 1000)

    return () => clearInterval(timer)
  }, [])

  const features = [
    {
      icon: <MapPin className="w-8 h-8 text-japanese-red" />,
      title: "GPS Temple Finder",
      description: "Find nearby temples and shrines based on your location",
      href: "/temples"
    },
    {
      icon: <Camera className="w-8 h-8 text-japanese-red" />,
      title: "Digital Collection",
      description: "Capture and organize your goshuin stamps digitally",
      href: "/collection"
    },
    {
      icon: <BookOpen className="w-8 h-8 text-japanese-red" />,
      title: "Cultural Guide",
      description: "Learn about goshuin etiquette and temple culture",
      href: "/guide"
    },
    {
      icon: <Globe className="w-8 h-8 text-japanese-red" />,
      title: "Official Information",
      description: "Access official temple websites and social media",
      href: "/temples"
    }
  ]

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative h-screen flex items-center justify-center overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-japanese-red/10 to-japanese-gold/10"></div>
        <div className="relative z-10 text-center px-4">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            <h1 className="text-5xl md:text-7xl font-bold text-japanese-black mb-6">
              Goshuin App
            </h1>
            <p className="text-xl md:text-2xl text-gray-600 mb-8 max-w-2xl mx-auto">
              Your Digital Cultural Passport to Japanese Temples and Shrines
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link href="/temples">
                <button className="bg-japanese-red text-white px-8 py-3 rounded-lg font-semibold hover:bg-red-700 transition-colors">
                  Explore Temples
                </button>
              </Link>
              <Link href="/guide">
                <button className="border-2 border-japanese-red text-japanese-red px-8 py-3 rounded-lg font-semibold hover:bg-japanese-red hover:text-white transition-colors">
                  Learn More
                </button>
              </Link>
            </div>
          </motion.div>
        </div>
        
        {/* Floating elements */}
        <motion.div
          animate={{ y: [0, -20, 0] }}
          transition={{ duration: 3, repeat: Infinity }}
          className="absolute top-20 left-10 text-japanese-red opacity-20"
        >
          ⛩️
        </motion.div>
        <motion.div
          animate={{ y: [0, 20, 0] }}
          transition={{ duration: 4, repeat: Infinity }}
          className="absolute bottom-20 right-10 text-japanese-gold opacity-20"
        >
          🏮
        </motion.div>
      </section>

      {/* Features Section */}
      <section className="py-20 px-4 bg-white">
        <div className="max-w-6xl mx-auto">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            className="text-center mb-16"
          >
            <h2 className="text-4xl font-bold text-japanese-black mb-4">
              Discover Japanese Culture
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Experience the beauty of Japanese temples and shrines through the ancient tradition of goshuin collection
            </p>
          </motion.div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {features.map((feature, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.6, delay: index * 0.1 }}
                className="group"
              >
                <Link href={feature.href}>
                  <div className="bg-white p-6 rounded-xl shadow-lg card-hover border border-gray-100">
                    <div className="mb-4 group-hover:scale-110 transition-transform">
                      {feature.icon}
                    </div>
                    <h3 className="text-xl font-semibold text-japanese-black mb-2">
                      {feature.title}
                    </h3>
                    <p className="text-gray-600">
                      {feature.description}
                    </p>
                  </div>
                </Link>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      {/* About Section */}
      <section className="py-20 px-4 bg-gray-50">
        <div className="max-w-4xl mx-auto text-center">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            <h2 className="text-4xl font-bold text-japanese-black mb-6">
              What is Goshuin?
            </h2>
            <p className="text-lg text-gray-600 mb-8 leading-relaxed">
              Goshuin (御朱印) are special stamps or calligraphy that you can receive at Japanese temples and shrines. 
              They serve as proof of your visit and are considered sacred items. Each goshuin is unique and beautifully 
              hand-written by temple priests, making them treasured souvenirs of your spiritual journey through Japan.
            </p>
            <div className="bg-white p-8 rounded-xl shadow-lg">
              <p className="text-sm text-gray-500 mb-2">
                Current time in Japan
              </p>
              <p className="text-2xl font-mono text-japanese-red">
                {currentTime.toLocaleTimeString('en-US', { 
                  timeZone: 'Asia/Tokyo',
                  hour12: false 
                })}
              </p>
            </div>
          </motion.div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-japanese-black text-white py-12 px-4">
        <div className="max-w-6xl mx-auto text-center">
          <h3 className="text-2xl font-bold mb-4">Goshuin App</h3>
          <p className="text-gray-300 mb-6">
            Your digital companion for exploring Japanese temple culture
          </p>
          <div className="flex justify-center space-x-6">
            <Link href="/guide" className="text-gray-300 hover:text-white transition-colors">
              Guide
            </Link>
            <Link href="/temples" className="text-gray-300 hover:text-white transition-colors">
              Temples
            </Link>
            <Link href="/collection" className="text-gray-300 hover:text-white transition-colors">
              Collection
            </Link>
          </div>
          <p className="text-gray-400 text-sm mt-8">
            © 2024 Goshuin App. All rights reserved.
          </p>
        </div>
      </footer>
    </div>
  )
}
