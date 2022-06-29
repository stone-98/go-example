package _interface

import (
	"io"
	"time"
)

// Artifact 人工制品
type Artifact interface {
	// Title 标题
	Title() string
	// Creators 作者列表
	Creators() []string
	// Created 创作日期
	Created() time.Time
}

// Text 文本
type Text interface {
	Pages() int
	Works() int
	PageSize() int
}

// Audio 音乐
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

// Video 视频
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

// Streamer 流
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
