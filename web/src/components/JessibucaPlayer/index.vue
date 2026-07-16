<template>
  <div class="jessibuca-player">
    <div class="player-toolbar">
      <el-select v-model="currentProtocol" size="small" style="width: 140px" @change="handleProtocolChange">
        <el-option v-for="item in availableProtocols" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <el-tag v-if="playing" type="success" size="small" effect="dark">播放中</el-tag>
      <el-tag v-else-if="errorMsg" type="danger" size="small" effect="dark">{{ errorMsg }}</el-tag>
      <el-tag v-else type="info" size="small" effect="dark">未播放</el-tag>
    </div>
    <div ref="containerRef" class="player-container" @dblclick="toggleFullscreen"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  urls: {
    type: Object,
    default: () => ({})
  },
  hasAudio: {
    type: Boolean,
    default: true
  }
})

const containerRef = ref(null)
const currentProtocol = ref('ws_flv')
const playing = ref(false)
const errorMsg = ref('')

let player = null

function loadJessibuca() {
  return new Promise((resolve, reject) => {
    if (window.Jessibuca) {
      resolve(window.Jessibuca)
      return
    }
    const script = document.createElement('script')
    script.src = '/jessibuca/jessibuca.min.js'
    script.onload = () => resolve(window.Jessibuca)
    script.onerror = () => reject(new Error('Failed to load jessibuca'))
    document.head.appendChild(script)
  })
}

const protocolMap = {
  ws_flv: { label: 'WS-FLV', key: 'ws_flv', getValue: (u) => u.ws_flv?.ws },
  http_flv: { label: 'HTTP-FLV', key: 'flv', getValue: (u) => u.flv?.http },
  hls: { label: 'HLS', key: 'hls', getValue: (u) => u.hls },
  rtsp: { label: 'RTSP', key: 'rtsp', getValue: (u) => u.rtsp },
  rtmp: { label: 'RTMP', key: 'rtmp', getValue: (u) => u.rtmp }
}

const availableProtocols = computed(() => {
  return Object.entries(protocolMap)
    .filter(([, v]) => v.getValue(props.urls))
    .map(([, v]) => ({ label: v.label, value: v.key }))
})

function getStreamUrl(protocol) {
  const def = protocolMap[protocol]
  return def ? def.getValue(props.urls) : null
}

function destroyPlayer() {
  if (player) {
    try {
      player.destroy()
    } catch {}
    player = null
  }
  playing.value = false
}

async function createPlayer() {
  destroyPlayer()
  if (!containerRef.value) return

  let JessibucaClass
  try {
    JessibucaClass = await loadJessibuca()
  } catch (err) {
    errorMsg.value = '播放器加载失败'
    console.error('Failed to load Jessibuca:', err)
    return
  }

  // Configuration matching wvp-GB28181-pro's jessibuca.vue for reliable playback
  player = new JessibucaClass({
    container: containerRef.value,
    videoBuffer: 0,                    // Minimum latency
    decoder: '/jessibuca/decoder.js',
    isResize: true,
    isFlv: false,                      // Auto-detect format (not forced FLV)
    isHls: false,
    useMSE: false,                     // Use WASM rendering, not MSE
    useWCS: false,                     // Use WASM rendering, not WebCodecs
    forceNoOffscreen: true,            // Keep rendering when tab in background
    isNotMute: true,                   // Unmute audio by default
    timeout: 10,                       // Connection timeout (seconds)
    loadingTimeout: 10,                // Loading timeout
    loadingTimeoutReplay: true,        // Auto-retry on loading timeout
    loadingTimeoutReplayTimes: 3,
    heartTimeout: 5,                   // Heartbeat check interval
    heartTimeoutReplay: true,          // Auto-reconnect on heartbeat loss
    heartTimeoutReplayTimes: 3,
    wasmDecodeErrorReplay: true,       // Retry on WASM decode errors
    keepScreenOn: true,                // Prevent screen sleep during playback
    loadingText: '请稍等, 视频加载中...',
    controlAutoHide: false,            // Keep controls visible
    debug: false,
    hotKey: true,
    showBandwidth: false,
    operateBtns: {
      fullscreen: false,
      screenshot: false,
      play: false,
      audio: false,
      recorder: false
    },
    supportDblclickFullscreen: false,
    useWebFullSreen: true,
    hiddenAutoPause: false,
    isFullResize: false,
    openWebglAlignment: false,
    wcsUseVideoRendcer: true,
    hasAudio: props.hasAudio
  })

  player.on('play', () => {
    playing.value = true
    errorMsg.value = ''
  })

  player.on('pause', () => {
    playing.value = false
  })

  player.on('error', (err) => {
    console.error('Jessibuca error:', err)
  })

  player.on('timeout', () => {
    console.warn('Jessibuca timeout, reconnecting...')
  })

  player.on('loadingTimeout', () => {
    console.warn('Jessibuca loading timeout, retrying...')
  })
}

async function startPlay() {
  if (!player) await createPlayer()
  const url = getStreamUrl(currentProtocol.value)
  if (!url) {
    errorMsg.value = '无可用流地址'
    return
  }
  errorMsg.value = ''
  player.play(url)
}

function handleProtocolChange() {
  startPlay()
}

function toggleFullscreen() {
  if (!containerRef.value) return
  if (document.fullscreenElement) {
    document.exitFullscreen()
  } else {
    containerRef.value.requestFullscreen()
  }
}

function stop() {
  destroyPlayer()
  errorMsg.value = ''
}

async function play() {
  await createPlayer()
  startPlay()
}

watch(() => props.urls, (newUrls) => {
  if (newUrls && Object.keys(newUrls).length > 0) {
    play()
  }
}, { deep: true })

onMounted(() => {
  if (props.urls && Object.keys(props.urls).length > 0) {
    play()
  }
})

onBeforeUnmount(() => {
  destroyPlayer()
})

defineExpose({ play, stop })
</script>

<style scoped>
.jessibuca-player {
  width: 100%;
  display: flex;
  flex-direction: column;
}

.player-toolbar {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 0;
}

.player-container {
  width: 100%;
  height: 480px;
  background: #000;
  border-radius: 4px;
  overflow: hidden;
}
</style>
