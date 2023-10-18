'use client'

export function convertToAlphanumericWithUnderscore(inputStr: string) {
    return inputStr.replace(/[^a-zA-Z0-9_]/g, '');
}