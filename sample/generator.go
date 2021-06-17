package sample

import (
  "github.com/peterliao96/PCBook/pb"
  "github.com/golang/protobuf/ptypes"
)


// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard{
  keyboard := &pb.Keyboard{
    Layout: randomKeyboardLayout(),
    Backlit: randomBool(),
  }
  return keyboard
}

func NewCPU() *pb.CPU{
  brand := randomCPUBrand()
  name := randomCPUName()
  numberCores := randomInt(2,8)
  numberThreads := randomInt(numberCores,12)

  minGhz := randomFloat64(2.0,3.5)
  maxGhz := randomFloat64(minGhz,5.0)
  cpu := &pb.CPU{
    Brand : brand,
    Name: name,
    NumberCores: uint32(numberCores), // type conversion
    NumberThreads: uint32(numberThreads),
    MinGhz: minGhz,
    MaxGhz: maxGhz,
  }
  return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU{
  brand := randomGPUBrand()
  name := randomGPUName(brand)

  minGhz := randomFloat64(1.0,1.5)
  maxGhz := randomFloat64(minGhz,2.0)

  memory := &pb.Memory{
    Value: uint64(randomInt(2,6)),
    Unit: pb.Memory_GIGABYTE,
  }
  gpu := &pb.GPU{
    Brand: brand,
    Name: name,
    MinGhz: minGhz,
    MaxGhz: maxGhz,
    Memory: memory,
  }
  return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *pb.Memory{
  ram := &pb.Memory{
    Value: uint64(randomInt(4,64)),
    Unit: pb.Memory_GIGABYTE,
  }
  return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *pb.Storage{
  ssd := &pb.Storage{
    Driver: pb.Storage_SSD,
    Memory: &pb.Memory{
      Value: uint64(randomInt(128,1024)),
      Unit: pb.Memory_GIGABYTE,
    },
  }
  return ssd
}

// NewHDD returns a new sample HDD
func NewSSD() *pb.Storage{
  hdd := &pb.Storage{
    Driver: pb.Storage_HDD,
    Memory: &pb.Memory{
      Value: uint64(randomInt(1,6)),
      Unit: pb.Memory_TERABYTE,
    },
  }
  return hdd
}

// NewScreen returns a new sample screen
func NewScreen() *pb.Screen{
  height := randomInt(1080,4320)
  width := height * 16/9
  screen := &pb.Screen{
    SizeInch: randomFloat32(13,17),
    Resolution: &pb.Screen_Resolution{
      Height: uint32(height),
      Width: uint32(width),
    },
    Panel: randomScreenPanel(),
    Multitouch: randomBool(),
  }
  return screen
}

// NewLaptop returns a new sample laptop
func NewLaptop()*pb.Laptop{
  brand := randomLaptopBrand()
  name := randomLaptopName(brand)
  laptop := &pb.Laptop{
    ID: randomID(),
    Brand: brand,
    Name: name,
    Cpu: NewCPU(),
    Ram: NewRAM(),
    Gpus: []*pb.GPU{NewGPU()},
    Storages: []*pb.Storage{NewSSD(),NewHDD()},
    Screen: NewScreen(),
    Keyboard: NewKeyboard(),
    Weight: &pb.Laptop_WeightKg{
      WeightKg: randomFloat64(1.0,3.0)
    },
    PriceUsd: randomFloat64(1500,3000),
    releaseYear: uint32(randomInt(2015,2019)),
    UpdatedAt: ptypes.TimestampNow(),
  }
}
