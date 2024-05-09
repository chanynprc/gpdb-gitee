{% set cgconfig_source = salt['pillar.get']('cgconfig:source_location', 'salt://files/cgconfig.conf.tpl') %}

set_cgroups:
  file.managed:
    - name: /etc/cgconfig.conf
    - source: {{ cgconfig_source }}