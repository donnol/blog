---
author: "jdlau"
date: 2023-10-17
linktitle: Wezterm极简配置文件
menu:
next:
prev:
title: Wezterm极简配置文件
weight: 10
categories: ['Rust', 'Wezterm']
tags: ['mark']
---

`Wezterm`极简配置文件

```lua
-- Pull in the wezterm API
local wezterm = require 'wezterm'

-- This table will hold the configuration.
local config = {}

-- In newer versions of wezterm, use the config_builder which will
-- help provide clearer error messages
if wezterm.config_builder then
  config = wezterm.config_builder()
end

-- This is where you actually apply your config choices

-- For example, changing the color scheme:
config.color_scheme = 'AdventureTime'
-- config.color_scheme = 'Batman'

config.enable_tab_bar = true
config.hide_tab_bar_if_only_one_tab = false
config.show_tab_index_in_tab_bar = false
config.tab_bar_at_bottom = true
config.tab_max_width = 25

-- Font
-- config.font = wezterm.font_with_fallback { 'JetBrains Mono' }
config.font_size = 14
config.freetype_load_target = "Light" -- Possible alternatives are `HorizontalLcd`, `Light`, `Mono`, `Normal`, `VerticalLcd`.

config.mouse_bindings = {  -- Paste on right-click
    {
        event = {
            Down = {
                streak = 1,
                button = 'Right'
            }
        },
        mods = 'NONE',
        action = wezterm.action {
            PasteFrom = 'Clipboard'
        }
    }, -- Change the default click behavior so that it only selects
    {
        event = {
            Up = {
                streak = 1,
                button = 'Left'
             }
         },
        mods = 'NONE',
        action = wezterm.action {
            CompleteSelection = 'PrimarySelection'
         }
    }, -- CTRL-Click open hyperlinks
    {
        event = {
            Up = {
                streak = 1,
                button = 'Left'
             }
         },
        mods = 'CTRL',
        action = 'OpenLinkAtMouseCursor'
    }
}

config.keys = {
    -- 使用 Ctrl+n 来新建标签页，而不是 Ctrl+Shift+T
    { key = "n",          mods = "ALT", action = wezterm.action { SpawnTab = "DefaultDomain" } },
    -- 使用 Alt+左右箭头来切换标签页
    { key = "LeftArrow",  mods = "ALT", action = wezterm.action { ActivateTabRelative = -1 } },
    { key = "RightArrow", mods = "ALT", action = wezterm.action { ActivateTabRelative = 1 } },
}

-- 定义 Alt+数字 切换到对应标签页的快捷键
for i = 1, 9 do
    table.insert(config.keys, {
        key = tostring(i),
        mods = 'ALT',
        action = wezterm.action.ActivateTab(i - 1),
    })
end

-- and finally, return the configuration to wezterm
return config

```
