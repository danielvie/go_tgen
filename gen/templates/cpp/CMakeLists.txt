
cmake_minimum_required(VERSION 3.20.0)
project(main)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_STANDARD 23)

add_library(mathlib SHARED
  lib/mathlib.cpp
)

target_include_directories(mathlib PRIVATE
  include
)

file(GLOB sources "src/*.cpp")
add_executable(main ${sources})
