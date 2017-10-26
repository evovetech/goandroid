package tech.evove.goandroid;

import core.Core;
import io.reactivex.Scheduler;

public class GoScheduler extends Scheduler {
    @Override
    public Worker createWorker() {
        // TODO: pooling
        return new GoWorker(Core.newWorker());
    }
}
