
var QQV = (rule, value, callback) => {

    if (value === '') {
        callback(new Error('输入QQ号'))
    } else if (/^[1-9][0-9]{4,10}$/.test(value)) {
        callback()
    } else {
        callback(new Error('输入正确的QQ号'))
    }
}

// 类似金钱,首位不为0,最多2位小数
export function checkNumPot2(rule, value, callback) {
    const reg = /(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)/
    if (!value) {
        return callback(new Error('请填写数字'))
    } else if (!reg.test(value)) {
        return callback(new Error('请填写数字,最多2位小数'))
    } else {
        callback()
    }
}

// 身份证
export function checkIdNum(rule, value, callback) {
    const reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/
    if (!value) {
        return callback(new Error('证件号码不能为空'))
    } else if (!reg.test(value)) {
        return callback(new Error('证件号码不正确'))
    } else {
        callback()
    }
}

// 整数
export function checkInterNum(rule, value, callback) {
    const reg = /^[0-9]*[1-9][0-9]*$/
    if (!value) {
        return callback(new Error('请填写整数'))
    } else if (!reg.test(value)) {
        return callback(new Error('请输入整数'))
    } else {
        callback()
    }
}

var UptextBox = (rule, value, callback) => {
    const reg=/^(-?\d+)(\.\d+)?$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须是数字'))
    }
}
var LefttextBox = (rule, value, callback) => {
    const reg=/^(-?\d+)(\.\d+)?$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须是数字'))
    }
}
var BottomtextBox = (rule, value, callback) => {
    const reg=/^(-?\d+)(\.\d+)?$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须是数字'))
    }
}
var RighttextBox = (rule, value, callback) => {
    const reg=/^(-?\d+)(\.\d+)?$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须是数字'))
    }
}

var zzs = (rule, value, callback) => {
    const reg=/^[1-9]\d*$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须为非零正整数'))
    }
}

var cityid = (rule, value, callback) => {
    const reg=/^(-?\d+)(\.\d+)?$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须是数字'))
    }
}

var en = (rule, value, callback) => {
    const reg=/[a-zA-Z]+$/
    if (value === '') {
        callback(new Error('不能为空'))
    } else if (reg.test(value)) {
        callback()
    } else {
        callback(new Error('必须英文'))
    }
}

export default {
    QQ: [{ required: true, validator: QQV, trigger: 'blur' }],
    phone: [{ required: true, pattern: /^1[34578]\d{9}$/, message: '目前只支持中国大陆的手机号码', trigger: 'blur' }],
    numPot2: [{ required: true, validator: checkNumPot2, trigger: 'blur' }],
    InterNum: [{ required: true, validator: checkInterNum, trigger: 'blur' }],
    name:[{ required: true, type: 'string',message:"名称不能为空", trigger: 'blur' }],
    file:[{ required: true, type: 'string',message:"文件地址不能为空", trigger: ['blur','change'] }],
    UptextBox:{ required: true, validator:UptextBox, trigger: ['blur','change'] },
    LefttextBox:{ required: true, validator:UptextBox, trigger: ['blur','change'] },
    BottomtextBox:{ required: true, validator:UptextBox, trigger: ['blur','change'] },
    RighttextBox:{ required: true, validator:UptextBox, trigger: ['blur','change'] },
    city:[{ required: true, type: 'array',message:"城市不能为空", trigger:['blur','change'] }],
    cityid:{ required: true, validator:cityid, trigger: ['blur','change'] },
    select:[{ required: true, type: 'string',message:"不能为空", trigger: ['blur','change']}],
    en:{ required: true, validator:en, trigger: ['blur','change'] },
    zzs:{ required: true, validator:zzs, trigger: ['blur','change'] },
}

