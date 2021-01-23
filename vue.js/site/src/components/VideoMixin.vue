<script lang='ts'>
import Vue from 'vue';
import { Component, Mixin, Mixins } from 'vue-mixin-decorator';
import jQuery from 'jQuery';
import axios from 'axios';
import Hls from 'hls.js';
import { VideoInfo } from '../model/models';

@Mixin
export default class VideoMixin extends Vue {
    public stop(): void {
        const isPlaying = this.$store.getters.isVideoPlaying;
        if (!isPlaying) {
            return;
        }

        (document.getElementById('project-video') as HTMLVideoElement).pause();

        this.$store.commit('setIsPlaying', {
            isVideoPlaying: false,
        });
        this.$store.commit('setIsDisplay', {
            isVideoDisplay: false,
        });

        $('.content, #sub-menu').css({
            opacity: 1,
        });
        $('#project-video').css({
            'opacity': 0,
            'z-index': -10,
        });
    }

    private async play(id: number): Promise<void> {
        const isDisplay = this.$store.getters.isVideoDisplay;
        if ( isDisplay ) {
            this.stop();
            return;
        }

        ($('#loading-modal') as any).modal({
            closable: false,
        }).modal('show');

        setTimeout(() => {
            const video = document.getElementById('project-video') as HTMLVideoElement;
            let result = false;

            // hls.js（HTTP Live Streaming）を利用し、サーバから動画データをダウンロードしながら再生
            if (Hls.isSupported()) {
                const hls = new Hls();
                hls.loadSource(`https://dggt9dial88dg.cloudfront.net/hls/video-${id}/index.m3u8`);
                hls.attachMedia(video);
                hls.on(Hls.Events.MANIFEST_PARSED, () => {
                    result = true;
                    video.play();

                    this.$store.commit('setIsDisplay', {
                        isVideoDisplay: true,
                    });
                    this.$store.commit('setIsPlaying', {
                        isVideoPlaying: true,
                    });

                    ($('.modal') as any).modal('hide');

                    $('.content, #sub-menu').css({
                        opacity: 0,
                    });

                    $('#project-video').css({
                        'opacity': 1,
                        'z-index': 10,
                    });
                });

                hls.on(Hls.Events.ERROR, (event, data) => {
                    ($('.modal') as any).modal('hide');
                    console.log(data);
                    if (data.details === 'manifestLoadError') {
                        alert('videoがアップロードされていません');
                    }
                });
            }
        }, 1000);
    }
}
</script>