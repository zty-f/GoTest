import { GroupSuccessInfo, getUserInfo } from '../api/';
import dayjs from 'dayjs';

const defaultSize = {
  w: 750,
  h: 1240,
  x: 370,
  y: 550,
};

// 纯业务代码 我也很无奈😮‍💨
export async function getPosterImgData(
  qrUrl: string,
  poster: string,
  info?: {
    avatar?: string;
    name?: string;
    id?: number;
    has_open_formal_course?: boolean;
    isNovice?: boolean;
  },
  codeSize?: {
    w: number;
    h: number;
  },
  size = defaultSize,
  count?: number,
  name?: string
) {
  const canvas = document.createElement('canvas');
  const { w, h, x, y } = size;
  canvas.width = w;
  canvas.height = h;
  canvas.style.backgroundColor = 'red';
  const cxt = canvas.getContext('2d')!;
  // cxt!.scale(2, 1);
  cxt.beginPath();
  // 背景图
  const bgImg = await initImg(poster);
  cxt.drawImage(bgImg as CanvasImageSource, 0, 0, w, h);
  const result = await getUserInfo();
  console.log(info?.isNovice, name);
  if ((name && info?.has_open_formal_course) || (name && !info?.isNovice)) {
    cxt.font = '24px 黑体 bolder';
    cxt.fillStyle = '#555555';
    cxt.fillText(name, 450, 386);
  } else if (name && !info?.has_open_formal_course) {
    cxt.font = '24px 黑体 bolder';
    cxt.fillStyle = '#555555';
    cxt.fillText(name, 450, 402);
  }
  if (
    (info?.avatar && info?.has_open_formal_course) ||
    (info?.avatar && !info?.isNovice)
  ) {
    const avatar = await initImg(info!.avatar);
    cxt.save();
    cxt.beginPath();
    cxt.arc(80, y + 406, 38, 0, 2 * Math.PI); //绘制圆圈
    cxt.clip(); //裁剪
    cxt.drawImage(avatar as CanvasImageSource, 40, y + 370, 80, 80); //定位在圆圈范围内便会出现
    cxt.restore();
    cxt.restore();
    cxt.font = '24px 黑体 bolder';
    cxt.fillStyle = 'rgba(255, 255, 255, 1)';
    cxt.fillText(info!.name!, 130, y + 426);
  } else if (info?.avatar && !info?.has_open_formal_course) {
    const avatar = await initImg(info!.avatar);
    cxt.save();
    cxt.beginPath();
    cxt.arc(80, y + 480, 38, 0, 2 * Math.PI); //绘制圆圈
    cxt.clip(); //裁剪
    cxt.drawImage(avatar as CanvasImageSource, 40, y + 444, 80, 80); //定位在圆圈范围内便会出现
    cxt.restore();
    cxt.restore();
    cxt.font = '24px 黑体 bolder';
    cxt.fillStyle = 'rgba(255, 255, 255, 1)';
    cxt.fillText(info!.name!, 130, y + 500);
  }
  const isNotice = count === 5;
  if (info?.has_open_formal_course || !info?.isNovice) {
    cxt.font = '26px 黑体 bolder';
    cxt.textAlign = 'center';
    cxt.fillStyle = 'rgba(51, 51, 51, 1)';
    cxt.fillText('已读天数', 140, 1070);
    cxt.font = '52px 黑体 bolder';
    cxt.fillStyle = '#333333';
    cxt.fillText(`${result.study_statistics.day_count}`, 140, 1130);
    const day = cxt.measureText(`${result.study_statistics.day_count}`);
    cxt.font = '26px 黑体 bolder';
    cxt.fillStyle = '#333333';
    cxt.fillText('天', 160 + parseInt(day.width / 2 + ''), 1130);
    cxt.font = '26px 黑体 bolder';
    cxt.fillStyle = 'rgba(51, 51, 51, 1)';
    cxt.fillText('已读本数', 360, 1070);
    cxt.font = '52px 黑体 bolder';
    cxt.fillStyle = '#333333';
    cxt.fillText(`${result.study_statistics.book_count}`, 360, 1130);
    const num = cxt.measureText(`${result.study_statistics.book_count}`);
    cxt.font = '26px 黑体 bolder';
    cxt.fillStyle = '#333333';
    cxt.fillText('本', 380 + parseInt(num.width / 2 + ''), 1130);
    cxt.font = '26px 黑体 bolder';
    cxt.fillStyle = 'rgba(51, 51, 51, 1)';
    cxt.fillText('累计阅读', 580, 1070);
    cxt.font = '52px 黑体 bolder';
    cxt.fillStyle = '#333333';
    const total_vocabulary = unitConverter(
      result.study_statistics.total_vocabulary
    );
    cxt.fillText(`${total_vocabulary}`, 580, 1130);
    const words = cxt.measureText(`${total_vocabulary}`);
    cxt.font = '26px 黑体 bolder';
    cxt.fillStyle = isNotice ? 'rgba(255, 255, 255, 1)' : '#333333';
    cxt.fillText(
      `${result.study_statistics.total_vocabulary < 10000 ? '词' : '万词'}`,
      result.study_statistics.total_vocabulary < 10000
        ? 600 + parseInt(words.width / 2 + '')
        : 610 + parseInt(words.width / 2 + ''),
      1130
    );
  }
  const qrImg = await initImg(qrUrl);
  if (codeSize) {
    cxt.drawImage(
      qrImg as CanvasImageSource,
      0,
      0,
      codeSize.w,
      codeSize.h,
      x + 175,
      y + 590,
      200,
      200
    );
  } else {
    cxt.drawImage(
      qrImg as CanvasImageSource,
      0,
      0,
      500,
      500,
      x + 145,
      y + 490,
      200,
      200
    );
  }

  return canvas.toDataURL('image/png');
}

async function initImg(url: string) {
  const img = new Image();
  if (!url.startsWith('data:image')) {
    img.setAttribute('crossOrigin', 'Anonymous');
  }
  img.src = url;
  return new Promise((resolve) => {
    img.onload = () => {
      resolve(img);
    };
  });
}

export function getDate(deadline: string) {
  const EndTime = new Date(deadline); //截止时间
  const NowTime = new Date();
  const t = EndTime.getTime() - NowTime.getTime();
  const d = Math.floor(t / 1000 / 60 / 60 / 24);
  const h = Math.floor((t / 1000 / 60 / 60) % 24);
  const m = Math.floor((t / 1000 / 60) % 60);
  const s = Math.floor((t / 1000) % 60);
  return {
    dd: d.toString().padStart(2, '0'),
    hh: h.toString().padStart(2, '0'),
    mm: m.toString().padStart(2, '0'),
    ss: s.toString().padStart(2, '0'),
  };
}

export function supplementList(list: GroupSuccessInfo) {
  if (list.members.length > 2) {
    list.members = list.members.splice(0, 2);
    return list;
  } else {
    for (let i = list.members.length; i < 2; i++) {
      list.members[i] = {
        member: {
          avatar: '',
          last_phone: '',
          user_id: 0,
        },
      };
    }
    return list;
  }
}

export function getTime(time: number) {
  const deadline = dayjs((time + 24 * 60 * 60) * 1000).format(
    'YYYY/MM/DD HH:mm:ss'
  );
  if ((time + 24 * 60 * 60) * 1000 - new Date().getTime() > 0) {
    return getDate(deadline);
  }
  return {
    dd: '00',
    hh: '00',
    mm: '00',
    ss: '00',
  };
}

function unitConverter(num: number) {
  if (!num || isNaN(num)) {
    return 0;
  }
  // 此处为防止字符串形式的数值进来，因为toFixed方法只能用于数值型数
  num = Number(num);
  if (Math.abs(num) > 100000000) {
    return (num / 100000000).toFixed(2);
  } else if (Math.abs(num) > 10000) {
    return (num / 10000).toFixed(2);
  } else {
    return num;
  }
}
