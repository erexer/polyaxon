# Generated by Django 3.0.9 on 2020-08-07 12:47

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ("coredb", "0001_initial"),
    ]

    operations = [
        migrations.RenameField(
            model_name="run", old_name="run_time", new_name="duration",
        ),
    ]
